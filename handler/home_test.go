package handler

import (
	"encoding/json"
	"fmt"
	"stock/tushare"
	"testing"
)

var jsonStr = `
{
	"002929": {
		"name": "N润建",
		"industry": "通信设备",
		"area": "广西",
		"pe": 31.8,
		"outstanding": 0.55,
		"totals": 2.21,
		"totalAssets": 238683.97,
		"liquidAssets": 227159.88,
		"fixedAssets": 6207.73,
		"reserved": 20553.17,
		"reservedPerShare": 0.93,
		"esp": 1.085,
		"bvps": 7.18,
		"pb": 4.8,
		"timeToMarket": 20180301,
		"undp": 70067.05,
		"perundp": 3.17,
		"rev": 21.16,
		"profit": 33.88,
		"gpr": 25.32,
		"npr": 8.62,
		"holders": 105228.0
	},
	"600901": {
		"name": "N苏租",
		"industry": "多元金融",
		"area": "江苏",
		"pe": 23.26,
		"outstanding": 6.4,
		"totals": 29.87,
		"totalAssets": 4737941.0,
		"liquidAssets": 0.0,
		"fixedAssets": 56044.36,
		"reserved": 103792.31,
		"reservedPerShare": 0.35,
		"esp": 0.29,
		"bvps": 2.59,
		"pb": 3.48,
		"timeToMarket": 20180301,
		"undp": 186222.8,
		"perundp": 0.62,
		"rev": 0.0,
		"profit": 0.0,
		"gpr": 0.0,
		"npr": 57.45,
		"holders": 575923.0
	}
}
`

func TestJson(t *testing.T) {
	basicStock := make(map[string]tushare.BasicStock)
	if err := json.Unmarshal([]byte(jsonStr), &basicStock); err != nil {
		t.Fatal(err)
	}
	fmt.Println(basicStock)
}
