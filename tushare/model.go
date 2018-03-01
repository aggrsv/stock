package tushare

type BasicStock struct {
	// 名称
	Name string
	// 所属行业
	Industry string
	// 地区
	Area string
	// 市盈率
	Pe string
	// 流通股本(亿)
	Outstanding string
	// 总股本(亿)
	Totals string
	// 总资产(万)
	TotalAssets string
	// 流动资产
	LiquidAssets string
	// 固定资产
	FixedAssets string
	// 公积金
	Reserved string
	// 每股公积金
	ReservedPerShare string
	// 每股收益
	Esp string
	// 每股净资
	Bvps string
	// 市净率
	Pb string
	// 上市日期
	TimeToMarket string
	// 未分利润
	Undp string
	// 每股未分配
	Perundp string
	// 收入同比(%)
	Rev string
	// 利润同比(%)
	Profit string
	// 毛利率(%)
	Gpr string
	// 净利润率(%)
	Npr string
	// 股东人数
	Holders string
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
