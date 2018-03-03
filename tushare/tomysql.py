#! /usr/bin/python
# coding: utf8

from sqlalchemy import create_engine
import tushare as ts

engine = create_engine('mysql://root:1234@192.168.1.240/tzstock?charset=utf8')

#存入数据库
#df.to_sql('basic_stock',engine)

print "正在导入基础数据"
#df = ts.get_stock_basics()
#df.to_sql('basic_stock',engine,if_exists='append')

# for year in range(2013,2018):
#     for season in range(1,5):
#             print "正在导入%s年第%s季度盈利能力\n"%(year,season)
#             df = ts.get_profit_data(year,season)
#             print "\n"
#             df.to_sql('profit_%s_%s'%(year,season),engine)


for year in range(2013,2018):
    for season in range(1,5):
            print "正在导入%s年第%s季度营运能力"%(year,season)
            df = ts.get_operation_data(year,season)
            print "\n"
            df.to_sql('operation_%s_%s'%(year,season),engine)


for year in range(2013,2018):
    for season in range(1,5):
            print "正在导入%s年第%s季度成长能力"%(year,season)
            df = ts.get_growth_data(year,season)
            print "\n"
            df.to_sql('growth_%s_%s'%(year,season),engine)

for year in range(2013,2018):
    for season in range(1,5):
            print "正在导入%s年第%s季度偿债能力"%(year,season)
            df = ts.get_debtpaying_data(year,season)
            print "\n"
            df.to_sql('debtpaying_%s_%s'%(year,season),engine)

for year in range(2013,2018):
    for season in range(1,5):
            print "正在导入%s年第%s季度现金流量"%(year,season)
            df = ts.get_debtpaying_data(year,season)
            print "\n"
            df.to_sql('cashflow_%s_%s'%(year,season),engine)

for year in range(2013,2018):
    for season in range(1,5):
            print "正在导入%s年第%s季度业绩报告"%(year,season)
            df = ts.get_report_data(year,season)
            print "\n"
            df.to_sql('creport_%s_%s'%(year,season),engine)