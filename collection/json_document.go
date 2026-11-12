package collection

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"unicode/utf8"
)

// JSONLimits bound parsing and the escaped, styled, wrapped body projection.
// Nonpositive values use defaults. Prefix content is owned by the caller.
type JSONLimits struct{ Bytes, Depth, Nodes, RenderBytes int }

func DefaultJSONLimits() JSONLimits {
	return JSONLimits{Bytes: DefaultBodyPreviewBytes, Depth: 64, Nodes: 4096, RenderBytes: 256 * 1024}
}

func (l JSONLimits) normalized() JSONLimits {
	d := DefaultJSONLimits()
	for _, pair := range []struct{ value, fallback *int }{{&l.Bytes, &d.Bytes}, {&l.Depth, &d.Depth}, {&l.Nodes, &d.Nodes}, {&l.RenderBytes, &d.RenderBytes}} {
		if *pair.value <= 0 {
			*pair.value = *pair.fallback
		}
	}
	// WHY: Decoder supports deeper documents, but the viewer's recursive walk
	// must retain a stack bound even when a caller raises configurable limits.
	l.Depth = min(l.Depth, 256)

	return l
}

// Nodes retain source spans rather than decoded values. Object children remain
// in source order; duplicate members have different preorder node identities.
type jsonNode struct {
	start, end, keyStart, keyEnd int
	parent, depth                int
	kind                         byte
	children                     []int
}

type jsonParser struct {
	source  []byte
	limits  JSONLimits
	decoder *json.Decoder
	nodes   []jsonNode
}

func parseJSON(source []byte, limits JSONLimits) ([]jsonNode, error) {
	if len(source) > limits.Bytes {
		return nil, fmt.Errorf("JSON folding byte limit (%d)", limits.Bytes)
	}

	if !utf8.Valid(source) || !json.Valid(source) {
		return nil, errors.New("invalid JSON: showing captured source")
	}

	decoder := json.NewDecoder(bytes.NewReader(source))
	decoder.UseNumber()

	p := jsonParser{source: source, limits: limits, decoder: decoder}
	if _, err := p.value(-1, 0, -1, 0); err != nil {
		return nil, err
	}

	return p.nodes, nil
}

func (p *jsonParser) start() int {
	position := int(p.decoder.InputOffset())
	for position < len(p.source) && bytes.ContainsRune([]byte(" \n\t\r,:"), rune(p.source[position])) {
		position++
	}

	return position
}

func (p *jsonParser) value(parent, depth, keyStart, keyEnd int) (int, error) {
	if depth > p.limits.Depth {
		return 0, fmt.Errorf("JSON folding depth limit (%d)", p.limits.Depth)
	}

	if len(p.nodes) >= p.limits.Nodes {
		return 0, fmt.Errorf("JSON folding node limit (%d)", p.limits.Nodes)
	}

	start := p.start()
	if depth >= p.limits.Depth && (p.source[start] == '{' || p.source[start] == '[') {
		return 0, fmt.Errorf("JSON folding depth limit (%d)", p.limits.Depth)
	}

	if _, err := p.decoder.Token(); err != nil {
		return 0, fmt.Errorf("JSON value: %w", err)
	}

	index := len(p.nodes)
	kind := p.source[start]

	p.nodes = append(p.nodes, jsonNode{start: start, parent: parent, depth: depth, keyStart: keyStart, keyEnd: keyEnd, kind: kind})
	if kind == '{' || kind == '[' {
		for p.decoder.More() {
			memberStart, memberEnd := -1, 0
			if kind == '{' {
				memberStart = p.start()
				if _, err := p.decoder.Token(); err != nil {
					return 0, fmt.Errorf("JSON member: %w", err)
				}

				memberEnd = int(p.decoder.InputOffset())
			}

			child, err := p.value(index, depth+1, memberStart, memberEnd)
			if err != nil {
				return 0, err
			}

			p.nodes[index].children = append(p.nodes[index].children, child)
		}

		if _, err := p.decoder.Token(); err != nil {
			return 0, fmt.Errorf("JSON container: %w", err)
		}
	}

	p.nodes[index].end = int(p.decoder.InputOffset())

	return index, nil
}
