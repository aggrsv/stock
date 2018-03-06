package qtimg

import (
	"fmt"
	"stock/entity"
	"stock/utils"
	"strings"

	"github.com/axgle/mahonia"
)

func Laster(symbol []string) ([]*entity.Stock, error) {
	var params []string
	if len(symbol) <= 0 {
		return nil, nil
	}
	for _, s := range symbol {
		params = append(params, entity.StockWithPrefix(s))
	}
	resp, err := utils.Get(fmt.Sprintf(entity.APILaster, strings.Join(params, ",")))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	enc := mahonia.NewDecoder("gbk")
	listStrs := strings.Split(string(resp), ";")
	stockList := make([]*entity.Stock, 0, len(listStrs))
	for _, stockStr := range listStrs {
		elements := strings.Split(enc.ConvertString(string(stockStr)), "~")
		if len(elements) < 49 {
			continue
		}
		stock := &entity.Stock{
			// 未知
			Unknow: elements[0],
			// 股票名称
			Name: elements[1],
			// 股票代码
			Symbol: elements[2],
			// 当前价格
			CurrentPrice: elements[3],
			// 昨收
			PrevClose: elements[4],
			// 今开
			Open: elements[5],
			// 成交量
			Volume: elements[6],
			// 外盘
			Sell: elements[7],
			// 内盘
			Buy: elements[8],
			// 买一
			BidPriceP1: elements[9],
			// 买一量
			BidVolumeP1: elements[10],
			// 买二
			BidPriceP2: elements[11],
			// 买二量
			BidVolumeP2: elements[12],
			// 买三
			BidPriceP3: elements[13],
			// 买三量
			BidVolumeP3: elements[14],
			// 买一
			BidPriceP4: elements[15],
			// 买一量
			BidVolumeP4: elements[16],
			// 买一
			BidPriceP5: elements[17],
			// 买一量
			BidVolumeP5: elements[18],
			// 卖一
			AskPriceP1: elements[19],
			// 卖一量
			AskVolumeP1: elements[20],
			// 卖二
			AskPriceP2: elements[21],
			// 卖二量
			AskVolumeP2: elements[22],
			// 卖三
			AskPriceP3: elements[23],
			// 卖三量
			AskVolumeP3: elements[24],
			// 卖四
			AskPriceP4: elements[25],
			// 卖四量
			AskVolumeP4: elements[26],
			// 卖五
			AskPriceP5: elements[27],
			// 卖五量
			AskVolumeP5: elements[28],
			// 最近逐笔成交
			CurrentDeal: elements[29],
			// 时间
			Date: elements[30],
			// 涨跌
			Change: elements[31],
			// 涨跌幅
			Chg: elements[32],
			// 最高
			High: elements[33],
			// 最低
			Low: elements[34],
			// 价格/成交量（手）/成交额
			PVA: elements[35],
			// 成交量（手）
			VolumeX: elements[36],
			// 成交额（万）
			Amount: elements[37],
			// 换手率
			Exchange: elements[38],
			// 市盈率
			MarkWin: elements[39],
			//
			UForty: elements[40],
			//
			HighX: elements[41],
			//
			LowX: elements[42],
			// 振幅
			Amplitude: elements[43],
			// 流通市值
			Capitalization: elements[44],
			// 总市值
			Value: elements[45],
			// 市净率
			PB: elements[46],
			// 涨停价
			SurgedLimit: elements[47],
			// 跌停价
			DeclineLimit: elements[48],
		}
		stockList = append(stockList, stock)
	}

	return stockList, nil
}

func Laster2(symbol []string) (string, error) {
	var params []string
	if len(symbol) <= 0 {
		return "", nil
	}
	for _, s := range symbol {
		params = append(params, entity.StockWithPrefix(s))
	}
	resp, err := utils.Get(fmt.Sprintf(entity.APILaster, strings.Join(params, ",")))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	enc := mahonia.NewDecoder("gbk")
	//
	return enc.ConvertString(string(resp)), nil

}
