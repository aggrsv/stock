package entity

//
type Stock struct {
	// 未知
	Unknow string
	// 股票名称
	Name string
	// 股票代码
	Symbol string
	// 当前价格
	CurrentPrice string
	// 昨收
	PrevClose string
	// 今开
	Open string
	// 成交量
	Volume string
	// 外盘
	Sell string
	// 内盘
	Buy string
	// 买一
	BidPriceP1 string
	// 买一量
	BidVolumeP1 string
	// 买二
	BidPriceP2 string
	// 买二量
	BidVolumeP2 string
	// 买三
	BidPriceP3 string
	// 买三量
	BidVolumeP3 string
	// 买一
	BidPriceP4 string
	// 买一量
	BidVolumeP4 string
	// 买一
	BidPriceP5 string
	// 买一量
	BidVolumeP5 string
	// 卖一
	AskPriceP1 string
	// 卖一量
	AskVolumeP1 string
	// 卖二
	AskPriceP2 string
	// 卖二量
	AskVolumeP2 string
	// 卖三
	AskPriceP3 string
	// 卖三量
	AskVolumeP3 string
	// 卖四
	AskPriceP4 string
	// 卖四量
	AskVolumeP4 string
	// 卖五
	AskPriceP5 string
	// 卖五量
	AskVolumeP5 string
	// 最近逐笔成交
	CurrentDeal string
	// 时间
	Date string
	// 涨跌
	Change string
	// 涨跌幅
	Chg string
	// 最高
	High string
	// 最低
	Low string
	// 价格/成交量（手）/成交额
	PVA string
	// 成交量（手）
	VolumeX string
	// 成交额（万）
	Amount string
	// 换手率
	Exchange string
	// 市盈率
	MarkWin string
	//
	UForty string
	//
	HighX string
	//
	LowX string
	// 振幅
	Amplitude string
	// 流通市值
	Capitalization string
	// 总市值
	Value string
	// 市净率
	PB string
	// 涨停价
	SurgedLimit string
	// 跌停价
	DeclineLimit string
}

type Income struct {
	// 利润总额
	IncomeBeforeTax string
	// 净利润
	NetIncome string
}


