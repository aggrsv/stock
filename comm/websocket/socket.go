package websocket

import (
	"io"
	"stock/comm/http"

	"github.com/gorilla/websocket"
)

const (
	TextMessage = websocket.TextMessage
	IntMessage  = websocket.BinaryMessage
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(ctx *http.Context) (*websocket.Conn, error) {
	scop := ctx.Scope()
	return upgrader.Upgrade(scop.ResponseWriter(), scop.Request(), nil)
}

//
func Read(conn *websocket.Conn) error {
	return nil
}

//
func Write(conn *websocket.Conn, messageType int, r io.Reader) error {
	w, err := conn.NextWriter(messageType)
	if err != nil {
		return err
	}
	if _, err := io.Copy(w, r); err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}
