package handler

import (
	"encoding/json"
	"fmt"
	"stock/comm/http"
	"stock/tushare"

	_ "stock/models/http"
)

//
func HomeHandler(ctx *http.Context) error {
	res, err := tushare.Tushare()
	if err != nil {
		fmt.Println("tushare error", err)
	}
	basicStock := make(map[string]tushare.BasicStock)
	if err := json.Unmarshal([]byte(res), &basicStock); err != nil {
		return err
	}
	ctx.Reply(basicStock)
	return nil
}
