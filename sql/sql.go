package sql

import (
	"database/sql"
	"database/sql/driver"
	"io"
	"reflect"
	"time"

	"github.com/go-sql-driver/mysql"
	"shendu.com/etc"
	"shendu.com/log"
)

var (
	ErrNoRows = sql.ErrNoRows
)

type DB struct {
	*sql.DB

	rdbs    []*sql.DB
	rdbChan chan *sql.DB
}

func (db *DB) Close() error {
	Close(db.DB)
	for _, rdb := range db.rdbs {
		Close(rdb)
	}
	return nil
}

func Default(dsn string) (*DB, error) {
	return Open(etc.String("sql", "driver"), dsn)
}

func Open(driver, dsn string) (*DB, error) {
	slaves, weight, err := Slaves(dsn)
	if err != nil {
		return nil, err
	}
	rdbs := make([]*sql.DB, len(slaves))
	weightMap := make(map[int]int)
	for i, slave := range slaves {
		weightMap[i] = slave.Weight
		db, err := sql.Open(driver, slave.Domain)
		if err != nil {
			return nil, err
		}
		db.SetMaxOpenConns(slave.MaxOpenConns)
		db.SetMaxIdleConns(slave.MaxIdelConns)
		rdbs[i] = db
	}
	rdbChan := make(chan *sql.DB, weight)
	for i := 0; i < weight; i++ {
		index := i % len(weightMap)
		rdbChan <- rdbs[index]
		weightMap[index] = weightMap[index] - 1
		if weightMap[index] == 0 {
			delete(weightMap, index)
		}
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(int(etc.Int("sql", "max_open_conns")))
	db.SetMaxIdleConns(int(etc.Int("sql", "max_idel_conns")))
	if weight == 0 || len(slaves) == 0 {
		rdbChan = make(chan *sql.DB, 1)
		rdbChan <- db
	}
	return &DB{db, rdbs, rdbChan}, nil
}

func Close(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Warn(err)
	}
}

func Rollback(tx *Tx) {
	if err := tx.Rollback(); err != nil {
		log.Warn(err)
	}
}

// 以下两个函数除了返回值，代码是一致的。
// 如果提出一个公共部分，那么将会对NewError的Caller值有影响，
// 所以暂且留着

func (db *DB) Exec(sql string, affect int64, args ...interface{}) error {
	rs, err := db.DB.Exec(sql, args...)
	if err != nil {
		return NewError().WithCause(err)
	}
	if affect > 0 {
		affected, err := rs.RowsAffected()
		if err != nil {
			return NewError().WithCause(err)
		}
		if affected != affect {
			return NewError("invalid rows affected").WithData(affect, affected)
		}
	}
	return nil
}

func (db *DB) ExecX(sql string, affect int64, args ...interface{}) (int64, error) {
	rs, err := db.DB.Exec(sql, args...)
	if err != nil {
		return 0, NewError().WithCause(err)
	}
	if affect > 0 {
		affected, err := rs.RowsAffected()
		if err != nil {
			return 0, NewError().WithCause(err)
		}
		if affected != affect {
			return 0, NewError("invalid rows affected").WithData(affect, affected)
		}
	}
	return rs.LastInsertId()
}

func (db *DB) QueryRow(sql string, args ...interface{}) *Row {
	rdb := <-db.rdbChan
	db.rdbChan <- rdb
	return &Row{rdb.QueryRow(sql, args...)}
}

func (db *DB) Query(sql string, args ...interface{}) (*Rows, error) {
	rdb := <-db.rdbChan
	db.rdbChan <- rdb
	rows, err := rdb.Query(sql, args...)
	if err != nil {
		return nil, NewError().WithCause(err)
	}
	return &Rows{rows}, nil
}

func (db *DB) Begin() (*Tx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, NewError().WithCause(err)
	}
	return &Tx{tx}, nil
}

type Row struct {
	*sql.Row
}

func (r *Row) Scan(dest ...interface{}) error {
	nullScanner := newNullScanner(dest...)
	if err := r.Row.Scan(nullScanner.nullType...); err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return NewError().WithCause(err)
	}
	return nullScanner.pop()
}

type Rows struct {
	*sql.Rows
}

func (r *Rows) Scan(dest ...interface{}) error {
	nullScanner := newNullScanner(dest...)
	if err := r.Rows.Scan(nullScanner.nullType...); err != nil {
		return NewError().WithCause(err)
	}
	return nullScanner.pop()
}

func (r *Rows) Err() error {
	if err := r.Rows.Err(); err != nil {
		return NewError().WithCause(err)
	}
	return nil
}

type Tx struct {
	*sql.Tx
}

func (tx *Tx) Commit() error {
	if err := tx.Tx.Commit(); err != nil {
		return NewError().WithCause(err)
	}
	return nil
}

// 以下两个函数除了返回值，代码是一致的。
// 如果提出一个公共部分，那么将会对NewError的Caller值有影响，
// 所以暂且留着

func (tx *Tx) Exec(sql string, affect int64, args ...interface{}) error {
	rs, err := tx.Tx.Exec(sql, args...)
	if err != nil {
		return NewError().WithCause(err)
	}
	if affect > 0 {
		affected, err := rs.RowsAffected()
		if err != nil {
			return NewError().WithCause(err)
		}
		if affected != affect {
			return NewError("invalid rows affected").WithData(affect, affected)
		}
	}
	return nil
}

func (tx *Tx) ExecX(sql string, affect int64, args ...interface{}) (int64, error) {
	rs, err := tx.Tx.Exec(sql, args...)
	if err != nil {
		return 0, NewError().WithCause(err)
	}
	if affect > 0 {
		affected, err := rs.RowsAffected()
		if err != nil {
			return 0, NewError().WithCause(err)
		}
		if affected != affect {
			return 0, NewError("invalid rows affected").WithData(affect, affected)
		}
	}
	return rs.LastInsertId()
}

func (tx *Tx) QueryRow(sql string, args ...interface{}) *Row {
	return &Row{tx.Tx.QueryRow(sql, args...)}
}

func (tx *Tx) Query(sql string, args ...interface{}) (*Rows, error) {
	rows, err := tx.Tx.Query(sql, args...)
	if err != nil {
		return nil, NewError().WithCause(err)
	}
	return &Rows{rows}, nil
}

func (tx *Tx) Rollback() error {
	if err := tx.Tx.Rollback(); err != nil {
		return NewError().WithCause(err)
	}
	return nil
}

type nullScanner struct {
	dest     []interface{}
	nullType []interface{}
}

func newNullScanner(dest ...interface{}) *nullScanner {
	nullType := make([]interface{}, len(dest))
	for i, item := range dest {
		switch item.(type) {
		case *string:
			nullType[i] = &sql.NullString{}
		case *int64, *int32, *int16, *int8, *int,
			*uint64, *uint32, *uint16, *uint8, *uint:
			nullType[i] = &sql.NullInt64{}
		case *float64, *float32:
			nullType[i] = &sql.NullFloat64{}
		case *bool:
			nullType[i] = &sql.NullBool{}
		case *time.Time:
			nullType[i] = &mysql.NullTime{}
		default:
			switch reflect.ValueOf(item).Elem().Kind() {
			case reflect.Int32:
				nullType[i] = &sql.NullInt64{}
			default:
				log.With("item", reflect.TypeOf(item)).Fatal("Unsupported type")
			}
		}
	}
	return &nullScanner{dest: dest, nullType: nullType}
}

func (ns *nullScanner) pop() error {
	for i, item := range ns.nullType {
		val, err := item.(driver.Valuer).Value()
		if err != nil {
			return NewError().WithCause(err)
		}
		if val != nil {
			dv := reflect.Indirect(reflect.ValueOf(ns.dest[i]))
			switch d := ns.dest[i].(type) {
			case *string:
				*d = val.(string)
			case *bool:
				*d = val.(bool)
			case *int64, *int32, *int16, *int8, *int:
				dv.SetInt(val.(int64))
			case *uint64, *uint32, *uint16, *uint8, *uint:
				dv.SetUint(uint64(val.(int64)))
			case *float64, *float32:
				dv.SetFloat(val.(float64))
			case *time.Time:
				t := val.(time.Time)
				*d = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local)
			default:
				switch reflect.ValueOf(ns.dest[i]).Elem().Kind() {
				case reflect.Int32:
					dv.SetInt(val.(int64))
				default:
					return NewError("unsupported dest type")
				}
			}
		}
	}
	return nil
}
