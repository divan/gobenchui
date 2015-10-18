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
func wshandler(ws *websocket.Conn, pool *WSPool) {
	conn := pool.Register(ws)
	defer func() {
		fmt.Println("[DEBUG] Closing connection")
		pool.Deregister(conn)
	}()

	for {
		val, ok := <-conn.ch
		if !ok {
			return
		}

		var data WSData
		if status, ok := val.(BenchmarkRun); ok {
			data = WSData{
				Type:      "status",
				Status:    InProgress,
				Commit:    status.Commit,
				StartTime: status.StartTime,
			}
		} else if set, ok := val.(BenchmarkSet); ok {
			data = WSData{
				Type:   "result",
				Result: set,
				Status: InProgress,
			}
		}

		if err := sendJSON(ws, data); err != nil {
			return
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
		// skip silently, as it's normal when client disconnects
		return err
	}

	return nil
}

// WSConn represents single websocket connection.
type WSConn struct {
	id int64
	ws *websocket.Conn
	ch chan interface{}
}

// WSPool holds registered websocket connections.
type WSPool map[int64]*WSConn

// Register registers new websocket connection and creates new channel for it.
func (pool WSPool) Register(ws *websocket.Conn) *WSConn {
	ch := make(chan interface{})
	id := time.Now().UnixNano()
	wsConn := &WSConn{
		id: id,
		ws: ws,
		ch: ch,
	}

	pool[id] = wsConn

	return wsConn
}

// Deregister removes connection from pool.
func (pool WSPool) Deregister(conn *WSConn) {
	for id, c := range pool {
		if id == conn.id {
			c.ws.Close()
			close(c.ch)
			delete(pool, id)
			return
		}
	}
}
