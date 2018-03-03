package handler

import (
	"errors"
	"stock/analyze"
	pb "stock/analyze/proto"
	"stock/comm/http"
	"sync"
)

func CodeNameHandler(ctx *http.Context) error {
	symbol := ctx.FormString("code")
	if symbol == "" {
		return errors.New("bad request")
	}
	list, err := analyze.CodeNameList(ctx, symbol)
	if err != nil {
		return err
	}
	ctx.Reply(list)
	return nil
}

func GrowingCompareHandler(ctx *http.Context) error {
	code := ctx.FormString("code")
	if code == "" {
		return errors.New("bad request")
	}
	growths, err := analyze.GrowingCompare(ctx, code)
	if err != nil {
		return err
	}
	ctx.Reply(growths)
	return nil
}

func ProfitHandler(ctx *http.Context) error {
	code := ctx.FormString("code")
	if code == "" {
		return errors.New("bad request")
	}
	profits, err := analyze.ProfitCompare(ctx, code)
	if err != nil {
		return err
	}
	ctx.Reply(profits)
	return nil
}

type integrate struct {
	Profits []*pb.Profit
	Growths []*pb.Growth
}

func IntegrateHandler(ctx *http.Context) error {
	code := ctx.FormString("code")
	if code == "" {
		return errors.New("bad request")
	}
	var (
		wg  sync.WaitGroup
		itg = &integrate{}
		// profit = make([]*pb.Profit, 0)
		// growth = make([]*pb.Growth, 0)
	)

	wg.Add(1)
	go func() error {
		growths, err := analyze.GrowingCompare(ctx, code)
		if err != nil {
			return err
		}
		itg.Growths = growths
		wg.Done()
		return nil
	}()

	wg.Add(1)
	go func() error {
		profits, err := analyze.ProfitCompare(ctx, code)
		if err != nil {
			return err
		}
		itg.Profits = profits
		wg.Done()
		return nil
	}()

	wg.Wait()
	ctx.Reply(itg)
	return nil
}
