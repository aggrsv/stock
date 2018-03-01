package tushare

type BasicStock struct {
	// 名称
	Name string
	// 所属行业
	Industry string
	// 地区
	Area string
	// 市盈率
	Pe float32
	// 流通股本(亿)
	Outstanding float32
	// 总股本(亿)
	Totals float32
	// 总资产(万)
	TotalAssets float32
	// 流动资产
	LiquidAssets float32
	// 固定资产
	FixedAssets float32
	// 公积金
	Reserved float32
	// 每股公积金
	ReservedPerShare float32
	// 每股收益
	Esp float32
	// 每股净资
	Bvps float32
	// 市净率
	Pb float32
	// 上市日期
	TimeToMarket float32
	// 未分利润
	Undp float32
	// 每股未分配
	Perundp float32
	// 收入同比(%)
	Rev float32
	// 利润同比(%)
	Profit float32
	// 毛利率(%)
	Gpr float32
	// 净利润率(%)
	Npr float32
	// 股东人数
	Holders float32
}

// 业绩报告
type Report struct {
	Code       string //代码
	Name       string //名称
	Esp        string //每股收益
	EpsYoy     string //每股收益同比(%)
	Bvps       string //每股净资产
	Roe        string //净资产收益率(%)
	Rpcf       string //每股现金流量(元)
	NetProfits string //净利润(万元)
	ProfitsYoy string //净利润同比(%)
	Distrib    string //分配方案
	ReportDate string //发布日期
}

// 盈利能力
type Profit struct {
	Code            string //代码
	Name            string //名称
	Roe             string //净资产收益率(%)
	NetProfitRatio  string //净利率(%)
	GrossProfitRate string //毛利率(%)
	NetProfits      string //净利润(万元)
	Esp             string //每股收益
	BusinessIncome  string //营业收入(百万元)
	Bips            string //每股主营业务收入(元)
}

// 成长能力
type Growth struct {
	Code string //代码
	Name string //名称
	Mbrg string //主营业务收入增长率(%)
	Nprg string //净利润增长率(%)
	Nav  string //净资产增长率
	Targ string //总资产增长率
	Epsg string //每股收益增长率
	Seg  string //股东权益增长率
}
