package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yunge/sphinx"
)

var (
	dbConn       *sql.DB
	err          error
	SphinxClient *sphinx.Client
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
	dbConn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//用于设置最大打开的连接数，默认值为0表示不限制。
	dbConn.SetMaxOpenConns(2000)
	//用于设置闲置的连接数。
	dbConn.SetMaxIdleConns(1000)
	dbConn.Ping()

	//sphinx
	SphinxClient = sphinx.NewClient().SetServer("localhost", 0).SetConnectTimeout(5000)
	if err := SphinxClient.Error(); err != nil {
		panic(err)
	}
}

func CheckError(err error) error {
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}

func DBQueryRow(rows *sql.Rows) (interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	//定义输出的类型
	result := make(map[string]interface{})
	//这个是sql查询出来的字段
	values := make([]interface{}, count)
	//保存sql查询出来的对应的地址
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		//scansql查询出来的字段的地址
		rows.Scan(valuePtrs...)

		//开始循环columns
		for i, col := range columns {
			var v interface{}
			//值
			val := values[i]
			//判读值的类型（interface类型）如果是byte，则需要转换成字符串
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			//保存
			result[col] = v
		}
	}
	_, ok := result["id"]
	if !ok {
		return "", errors.New("params invalid")
	}
	return result, nil
}

func DBQueryRows(rows *sql.Rows) (interface{}, error) {
	if rows == nil {
		return "", errors.New("empty data")
	}
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	if len(tableData) == 0 {
		return "", errors.New("invalid params")
	}
	return tableData, nil
}
