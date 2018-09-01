package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	//mysql connecter
	_ "github.com/go-sql-driver/mysql"
	"github.com/yunge/sphinx"
)

var (
	dbConn       *sql.DB
	sphinxClient *sphinx.Client
	redisClient  *redis.Client

	err error
)

//Init initinal start
func Init() {
	mysqlConnect()
}

//mysqlConnect mysql客户端
func mysqlConnect() {
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
	dbConn.SetMaxIdleConns(0)
}

//RedisConnect redis客户端
func redisConnect() {
	redisHost := beego.AppConfig.String("redishost")
	redisport := beego.AppConfig.String("redisport")
	redisPassword := beego.AppConfig.String("redispassword")
	redisDB, err := beego.AppConfig.Int("redisdb")
	if err != nil {
		redisDB = 0
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisport,
		Password: redisPassword,
		DB:       redisDB,
	})
}

//SphinxConnect conection aboard
func sphinxConnect() {
	sphinxClient := sphinx.NewClient().SetServer("localhost", 0).SetConnectTimeout(1000)
	if err := sphinxClient.Error(); err != nil {
		panic(err)
	}
	sphinxClient.SetMatchMode(sphinx.SPH_MATCH_ANY)
	fields := map[string]int{"post_intro": 3, "post_content": 2, "post_title": 1, "post_author": 4}
	sphinxClient.SetFieldWeights(fields)
}

//CheckError check error
func CheckError(err error) error {
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

//TableName get db table name
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
