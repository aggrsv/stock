package qtimg

import (
	"fmt"
	"testing"
)

func TestLaster(t *testing.T) {
	stock, err := Laster([]string{"600519", "000001"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(stock)
}
