// Package workbench serves real Go stories in isolated pseudo-terminal sessions.
package workbench

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync/atomic"
	"time"

	"github.com/coder/websocket"
	"github.com/creack/pty"
	"github.com/lanej/charmed/story"
)

type Server struct {
	executable string
	mux        *http.ServeMux
	ctx        context.Context
	cancel     context.CancelFunc
	active     atomic.Int64
	slots      chan struct{}
}

func NewServer(executable string) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	s := &Server{executable: executable, mux: http.NewServeMux(), ctx: ctx, cancel: cancel, slots: make(chan struct{}, 16)}
	s.mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{"status": "ok", "activeSessions": s.active.Load()})
	})
	s.mux.HandleFunc("GET /api/stories", func(w http.ResponseWriter, r *http.Request) { writeJSON(w, story.Definitions()) })
	s.mux.HandleFunc("POST /api/frame", s.frame)
	s.mux.HandleFunc("GET /api/session", s.session)
	return s
}
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.mux.ServeHTTP(w, r) }
func (s *Server) Close()                                           { s.cancel() }
func (s *Server) ActiveSessions() int64                            { return s.active.Load() }
func writeJSON(w http.ResponseWriter, value any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(value)
}

type frameRequest struct {
	ID     string       `json:"id"`
	Config story.Config `json:"config"`
}

func (s *Server) frame(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 16<<10)
	request := frameRequest{Config: story.DefaultConfig()}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		http.Error(w, "expected one JSON object", http.StatusBadRequest)
		return
	}
	frame, err := story.Frame(request.ID, request.Config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, map[string]string{"frame": frame})
}

func DecodeConfig(encoded string) (story.Config, error) {
	c := story.DefaultConfig()
	if encoded == "" {
		return c, nil
	}
	if len(encoded) > 24<<10 {
		return c, errors.New("configuration too large")
	}
	data, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return c, err
	}
	if err = json.Unmarshal(data, &c); err != nil {
		return c, err
	}
	return c, c.Validate()
}

func (s *Server) session(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if _, err := story.Find(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config, err := DecodeConfig(r.URL.Query().Get("config"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	select {
	case s.slots <- struct{}{}:
		defer func() { <-s.slots }()
	default:
		http.Error(w, "all preview sessions are in use", http.StatusTooManyRequests)
		return
	}
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		return
	}
	defer conn.CloseNow()
	conn.SetReadLimit(64 << 10)
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	stop := context.AfterFunc(s.ctx, cancel)
	defer stop()
	temp, err := os.MkdirTemp("", "charmed-session-")
	if err != nil {
		return
	}
	defer os.RemoveAll(temp)
	eventsR, eventsW, err := os.Pipe()
	if err != nil {
		return
	}
	defer eventsR.Close()
	encoded, _ := json.Marshal(config)
	cmd := exec.CommandContext(ctx, s.executable, "session", id, "--config", base64.RawURLEncoding.EncodeToString(encoded))
	cmd.Env = append(os.Environ(), "TERM=xterm-256color", "COLORTERM=truecolor", "CHARMED_EVENTS_FD=3", "CHARMED_FIXTURE_DIR="+temp)
	cmd.ExtraFiles = []*os.File{eventsW}
	cmd.Cancel = func() error { return cmd.Process.Signal(os.Interrupt) }
	cmd.WaitDelay = 2 * time.Second
	terminal, err := pty.StartWithSize(cmd, &pty.Winsize{Cols: uint16(config.Columns), Rows: uint16(config.Rows)})
	eventsW.Close()
	if err != nil {
		_ = conn.Write(ctx, websocket.MessageText, []byte(`{"type":"error","message":"could not start story"}`))
		return
	}
	defer terminal.Close()
	s.active.Add(1)
	defer s.active.Add(-1)
	write := func(kind websocket.MessageType, data []byte) error {
		deadline, done := context.WithTimeout(ctx, 5*time.Second)
		defer done()
		return conn.Write(deadline, kind, data)
	}
	go func() {
		buffer := make([]byte, 16<<10)
		for {
			n, err := terminal.Read(buffer)
			if n > 0 {
				if write(websocket.MessageBinary, buffer[:n]) != nil {
					cancel()
					return
				}
			}
			if err != nil {
				return
			}
		}
	}()
	go func() {
		scanner := bufio.NewScanner(eventsR)
		scanner.Buffer(make([]byte, 4096), 64<<10)
		for scanner.Scan() {
			data, err := json.Marshal(struct {
				Type  string          `json:"type"`
				Event json.RawMessage `json:"event"`
			}{"event", append(json.RawMessage(nil), scanner.Bytes()...)})
			if err == nil {
				if write(websocket.MessageText, data) != nil {
					return
				}
			}
		}
	}()
	done := make(chan struct{})
	go func() { defer close(done); _ = cmd.Wait(); cancel() }()
	defer func() { cancel(); terminal.Close(); <-done }()
	for {
		kind, data, err := conn.Read(ctx)
		if err != nil {
			return
		}
		if kind != websocket.MessageText {
			return
		}
		var message struct {
			Type    string `json:"type"`
			Data    string `json:"data"`
			Columns int    `json:"columns"`
			Rows    int    `json:"rows"`
		}
		if json.Unmarshal(data, &message) != nil {
			return
		}
		switch message.Type {
		case "input":
			if _, err = io.WriteString(terminal, message.Data); err != nil {
				return
			}
		case "resize":
			if message.Columns < 20 || message.Columns > 240 || message.Rows < 8 || message.Rows > 80 {
				return
			}
			if err = pty.Setsize(terminal, &pty.Winsize{Cols: uint16(message.Columns), Rows: uint16(message.Rows)}); err != nil {
				return
			}
		case "step":
			if _, err = terminal.Write([]byte{14}); err != nil {
				return
			}
		default:
			return
		}
	}
}
