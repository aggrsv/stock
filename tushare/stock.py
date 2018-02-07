import tushare

def stock_basics(sybmol):
    basics =  tushare.get_stock_basics()
    return  basics.to_json(orient='index')

def report_data(year,season):
    report = tushare.get_report_data(year,season)
    return report.to_json(orient='index')
