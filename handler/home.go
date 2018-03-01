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
	basicStock := make([]*tushare.BasicStock, 0)
	if err := json.Unmarshal([]byte(res), &basicStock); err != nil {
		return err
	}

	ctx.Reply(basicStock)
	return nil
}

func ProfitHandler(ctx *http.Context) error {
	res := tushare.GetProfit("2017", "1")
	profit := make([]*tushare.Profit, 0)
	if err := json.Unmarshal([]byte(res), &profit); err != nil {
		return err
	}
	ctx.Reply(profit)
	return nil
}
