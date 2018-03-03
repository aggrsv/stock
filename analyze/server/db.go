package main

import (
	"fmt"
	pb "stock/analyze/proto"
	"stock/etc"
	"stock/sql"
	"time"
)

var (
	dsn = etc.String("sql", "analyze")
)

type db struct {
	*sql.DB
}

func NewDb() *db {
	d, err := sql.Default(dsn)
	if err != nil {

	}
	return &db{
		DB: d,
	}
}

func (db *db) Close() error {
	return db.DB.Close()
}

func (db *db) codeNameList(symbol string) ([]*pb.CodeName, error) {
	rows, err := db.Query(fmt.Sprintf(codenamesql, symbol, symbol))
	if err != nil {
		return nil, err
	}
	defer sql.Close(rows)
	cns := make([]*pb.CodeName, 0)
	for rows.Next() {
		codename := &pb.CodeName{}
		if err := rows.Scan(&codename.Name, &codename.Code); err != nil {
			return nil, err
		}
		cns = append(cns, codename)
	}
	return cns, nil
}

const codenamesql = `select name, code from basic_stock where code like '%%%s%%' or name like '%%%s%%'`

func (db *db) GrowingCompare(code string) ([]*pb.Growth, error) {
	year := time.Now().Year()
	growths := make([]*pb.Growth, 0)
	for i := 1; i <= 5; i++ {
		pyear := year - i
		growth := &pb.Growth{}
		if err := db.QueryRow(fmt.Sprintf(growingcompareSql, pyear, pyear, pyear, pyear, code), code).Scan(
			&growth.Code, &growth.Name, &growth.Mbrg, &growth.Nprg, &growth.Nav,
			&growth.Targ, &growth.Epsg, &growth.Seg,
		); err != nil {
			return nil, err
		}
		if growth.Code == "" {
			continue
		}
		growth.Year = int32(pyear)
		growths = append(growths, growth)
	}
	return growths, nil
}

const growingcompareSql = `
select main.code, main.name, TRUNCATE(avg(mbrg),2) , TRUNCATE(avg(nprg),2), TRUNCATE(avg(nav),2), TRUNCATE(avg(targ),2), TRUNCATE(avg(epsg),2), TRUNCATE(avg(seg),2)
from (
select *  from growth_%d_4  
UNION
select *  from growth_%d_3  
UNION
select *  from growth_%d_2 
UNION
select *  from growth_%d_1
) as main where main.code = ? or main.name like '%%%s%%'

`

func (db *db) ProfitCompare(code string) ([]*pb.Profit, error) {
	year := time.Now().Year()
	profits := make([]*pb.Profit, 0)
	for i := 1; i <= 5; i++ {
		pyear := year - i
		profit := &pb.Profit{}
		if err := db.QueryRow(fmt.Sprintf(profitcompareSql, pyear, pyear, pyear, pyear, code), code).Scan(
			&profit.Code, &profit.Name, &profit.Roe, &profit.NetProfitRatio, &profit.GrossProfitRate,
			&profit.NetProfits, &profit.Esp, &profit.BusinessIncome, &profit.Bips,
		); err != nil {
			return nil, err
		}
		if profit.Code == "" {
			continue
		}
		profit.Year = int32(pyear)
		profits = append(profits, profit)
	}
	return profits, nil
}

const profitcompareSql = `
select main.code, main.name, TRUNCATE(avg(roe),2), TRUNCATE(avg(net_profit_ratio),2),TRUNCATE(avg(gross_profit_rate),2),
	  TRUNCATE(avg(net_profits),2),TRUNCATE(avg(eps),2),TRUNCATE(avg(business_income),2),TRUNCATE(avg(bips),2)
from (
select *  from profit_%d_4  
UNION
select *  from profit_%d_3  
UNION
select *  from profit_%d_2 
UNION
select *  from profit_%d_1
) as main where main.code = ? or main.name like '%%%s%%'

`
