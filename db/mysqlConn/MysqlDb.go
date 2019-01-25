package mysqlConn

import (
	"database/sql"
	"fmt"

	_ "github.com/Go-SQL-Driver/Mysql"
)

type MysqlDb struct {
	dns     string
	dbptr   *sql.DB
	conn    *sql.Tx
	rows    *sql.Rows
	res     sql.Result
	err     error
	IsDebug bool
}

func (db *MysqlDb) checkErr(param ...interface{}) {
	if db.err != nil {
		fmt.Println("error: ", db.err, param) //deal error here
		panic(db.err)
	}
}

func (db *MysqlDb) IsConnected() {
	err := db.dbptr.Ping()
	if err != nil {
		fmt.Println("PING ERROR: ", err, err.Error())
		db.Close()
		db.Open(db.dns)
	}
}

func (db *MysqlDb) Open(dns string) {
	db.dns = dns
	// "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8"
	db.dbptr, db.err = sql.Open("mysql", dns)
	db.checkErr()
}

func (db *MysqlDb) Query(sqlstr string, args ...interface{}) *sql.Rows {
	db.IsConnected()

	if nil != db.rows {
		db.rows.Close()
	}

	if db.IsDebug {
		fmt.Println(sqlstr, args)
	}

	db.rows, db.err = db.dbptr.Query(sqlstr, args...)
	db.checkErr()

	return db.rows
}

func (db *MysqlDb) Query2(sqlstr string, args ...interface{}) *sql.Rows {
	db.IsConnected()

	if db.IsDebug {
		fmt.Println(sqlstr, args)
	}

	rows, err := db.dbptr.Query(sqlstr, args...)
	db.err = err
	db.checkErr()

	return rows
}

// update and insert
func (db *MysqlDb) Exec(sqlstr string, args ...interface{}) bool {
	db.IsConnected()

	if db.IsDebug {
		fmt.Println(sqlstr, args)
	}

	if nil == db.conn {
		db.res, db.err = db.dbptr.Exec(sqlstr, args...)
		db.checkErr()

		return true
	}

	db.res, db.err = db.conn.Exec(sqlstr, args...)
	db.checkErr()

	return true
}

func (db *MysqlDb) BeginTrans() {
	db.IsConnected()

	db.conn, db.err = db.dbptr.Begin()
	db.checkErr()
}

func (db *MysqlDb) Rollback() {
	if nil != db.conn {
		//fmt.Println("回退事务")
		db.conn.Rollback()
	}
}

func (db *MysqlDb) Commit() {
	if nil != db.conn {
		fmt.Println("提交事务")
		db.conn.Commit()
	}
}

func (db *MysqlDb) Close() {
	if nil != db.rows {
		fmt.Println("db.rows close.！")
		db.rows.Close()
	}

	if nil != db.conn {
		db.conn.Commit()
	}

	if nil != db.dbptr {
		defer db.dbptr.Close()
	}
}
