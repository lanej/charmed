package workbench_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/workbench"
)

func TestFrameEndpointValidatesAndRendersGo(t *testing.T) {
	preview := workbench.NewServer("unused")
	defer preview.Close()
	server := httptest.NewServer(preview)
	defer server.Close()
	for _, tc := range []struct {
		body   string
		status int
	}{
		{`{"id":"input/text"}`, 200},
		{`{"id":"not/a/story"}`, 400},
		{`{"id":"input/text","config":{"columns":10000}}`, 400},
		{`{"id":"input/text","shell":"anything"}`, 400},
		{`{"id":"input/text"} {}`, 400},
	} {
		response, err := http.Post(server.URL+"/api/frame", "application/json", strings.NewReader(tc.body))
		if err != nil {
			t.Fatal(err)
		}
		data, _ := io.ReadAll(response.Body)
		response.Body.Close()
		if response.StatusCode != tc.status {
			t.Fatalf("%s: status %d: %s", tc.body, response.StatusCode, data)
		}
		if tc.status == 200 && !bytes.Contains(data, []byte("Text input")) {
			t.Fatal("frame does not contain Go fixture output")
		}
	}
}

func TestNativeSessionInputResizeAndCleanup(t *testing.T) {
	binary := filepath.Join(t.TempDir(), "workbench")
	build := exec.Command("go", "build", "-o", binary, "../cmd/workbench")
	if data, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build story runner: %v\n%s", err, data)
	}
	preview := workbench.NewServer(binary)
	defer preview.Close()
	server := httptest.NewServer(preview)
	defer server.Close()
	c := story.DefaultConfig()
	c.Value = ""
	raw, _ := json.Marshal(c)
	url := "ws" + strings.TrimPrefix(server.URL, "http") + "/api/session?id=input/text&config=" + base64.RawURLEncoding.EncodeToString(raw)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.CloseNow()
	var output strings.Builder
	for !strings.Contains(output.String(), "Text input") {
		kind, data, err := conn.Read(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if kind == websocket.MessageBinary {
			output.Write(data)
		}
	}
	if err = conn.Write(ctx, websocket.MessageText, []byte(`{"type":"resize","columns":40,"rows":16}`)); err != nil {
		t.Fatal(err)
	}
	if err = conn.Write(ctx, websocket.MessageText, []byte(`{"type":"input","data":"q\r"}`)); err != nil {
		t.Fatal(err)
	}
	for {
		kind, data, err := conn.Read(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if kind == websocket.MessageText && bytes.Contains(data, []byte(`"value":"q"`)) {
			break
		}
	}
	if preview.ActiveSessions() != 1 {
		t.Fatal("live story process is not tracked")
	}
	conn.Close(websocket.StatusNormalClosure, "test complete")
	deadline := time.NewTimer(4 * time.Second)
	defer deadline.Stop()
	tick := time.NewTicker(20 * time.Millisecond)
	defer tick.Stop()
	for preview.ActiveSessions() != 0 {
		select {
		case <-deadline.C:
			t.Fatal("preview process was not reaped after socket close")
		case <-tick.C:
		}
	}
}
