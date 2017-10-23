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
	// orm.RegisterDataBase("default", "mysql", dsn)
	// orm.RegisterModel(new(Posts))
	db_conn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
