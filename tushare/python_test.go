package tushare

import "testing"

func TestTushare(t *testing.T) {
	if err := Tushare(); err != nil {
		t.Fatal(err)
	}
}
