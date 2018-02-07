package qtimg

import (
	"fmt"
	"testing"
)

func TestLaster(t *testing.T) {
	stock, err := Laster("sh600519")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(stock)
}
