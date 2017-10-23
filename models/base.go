package models

import (
	"database/sql"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db_conn *sql.DB
	err     error
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	db_conn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//用于设置最大打开的连接数，默认值为0表示不限制。
	db_conn.SetMaxOpenConns(2000)
	//用于设置闲置的连接数。
	db_conn.SetMaxIdleConns(1000)
	db_conn.Ping()
}

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
