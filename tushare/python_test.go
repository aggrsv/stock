package tushare

import (
	"fmt"
	"testing"
)

func TestTushare(t *testing.T) {
	res, err := Tushare()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("======", res)
}

func TestProfit(t *testing.T) {
	res, err := GetProfit("2017", "3")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
