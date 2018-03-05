package entity

import (
	"fmt"
	"strings"
)

// 600***, 601*** 上证A股;
// 000***, 002*** 深圳A股;
// 400*** 三板市场股票
// 300*** 创业板

func StockWithPrefix(code string) string {
	switch {
	case strings.HasPrefix(code, "000"):
		return fmt.Sprintf("sz%s", code)
	case strings.HasPrefix(code, "002"):
		return fmt.Sprintf("sz%s", code)
	case strings.HasPrefix(code, "600"):
		return fmt.Sprintf("sh%s", code)
	case strings.HasPrefix(code, "601"):
		return fmt.Sprintf("sh%s", code)
	}
	return code
}
