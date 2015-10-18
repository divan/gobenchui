package main

import (
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

// WSData used for WebSocket commincation with frontend.
type WSData struct {
	Type   string       `json:"type"` // "status" or "result"
	Result BenchmarkSet `json:"result,omitempty"`

	Status    Status    `json:"status,omitempty"`
	Progress  float64   `json:"progress,omitempty"`
	Commit    Commit    `json:"commit,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
}

// wshandler is a handler for websocket connection.
func wshandler(ws *websocket.Conn, resCh chan BenchmarkSet, runCh chan BenchmarkRun) {
	defer func() {
		fmt.Println("[DEBUG] Closing connection")
		ws.Close()
	}()
	for {
		select {
		case status, ok := <-runCh:
			if !ok {
				runCh = nil
				if resCh == nil {
					return
				}
			}
			data := WSData{
				Type:      "status",
				Status:    InProgress,
				Commit:    status.Commit,
				StartTime: status.StartTime,
			}

			if err := sendJSON(ws, data); err != nil {
				return
			}
		case set, ok := <-resCh:
			if !ok {
				resCh = nil
				if runCh == nil {
					return
				}
			}
			data := WSData{
				Type:   "result",
				Result: set,
				Status: InProgress,
			}

			if err := sendJSON(ws, data); err != nil {
				return
			}
		}
	}
}

// sendJSON is a wrapper for sending JSON encoded data to websocket
func sendJSON(ws *websocket.Conn, data interface{}) error {
	body, err := json.MarshalIndent(data, "  ", "    ")
	if err != nil {
		fmt.Println("[ERROR] JSON encoding failed", err)
		return err
	}

	_, err = ws.Write(body)
	if err != nil {
		fmt.Println("[ERROR] WebSocket send failed", err)
		return err
	}

	return nil
}
