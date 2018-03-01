import tushare

def stock_basics(sybmol):
    basics =  tushare.get_stock_basics()
    return  basics.to_json(orient='index')

# 业绩报告
def report(year,season):
    report = tushare.get_report_data(year,season)
    return report.to_json(orient='index')

# 盈利能力
def profit(year,season):
    profit = tushare.get_profit_data(year,season)
    return profit.to_json(orient='index')

#成长能力
def growth(year,season):
    growth = tushare.get_growth_data(year,season)
    return growth.to_json(orient='index')

