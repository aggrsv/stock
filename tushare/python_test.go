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
