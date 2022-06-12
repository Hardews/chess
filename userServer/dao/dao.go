package dao

import (
	"database/sql"
	"github.com/go-redis/redis"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dB *sql.DB
)

func InitDB() {
	var db, err = sql.Open("mysql", "root:lmh123@tcp(127.0.0.1:3306)/chess")
	if err != nil {
		log.Fatal(err)
	}

	dB = db
}

// 声明一个全局的rdb变量
var rdb *redis.Client

func InitClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}
