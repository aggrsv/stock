package handler

import (
	"bytes"
	"errors"
	"fmt"
	"stock/comm/http"
	"stock/comm/websocket"
	"stock/qtimg"
	"strings"
)

func RealTimePriceHandler(ctx *http.Context) error {
	code := ctx.FormString("code")
	if code == "" {
		return errors.New("bad request")
	}
	price, err := qtimg.Laster(strings.Split(code, ","))
	if err != nil {
		fmt.Println("get laster price error ==>:", err)
		return err
	}
	ctx.Reply(price)
	return nil
}

func RealTimePrice2Handler(ctx *http.Context) error {
	code := ctx.FormString("code")
	if code == "" {
		return errors.New("bad request")
	}
	conn, err := websocket.Upgrade(ctx)
	if err != nil {
		return err
	}
	for {
		price, err := qtimg.Laster2(strings.Split(code, ","))
		if err != nil {
			fmt.Println("get laster price error ==>:", err)
			return err
		}
		r := bytes.NewReader([]byte(price))
		if err := websocket.Write(conn, websocket.TextMessage, r); err != nil {
			return err
		}
	}
	return nil
}
