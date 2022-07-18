package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-redis/redis"
	"log"
)

var (
	dB *gorm.DB
)

func InitDB() {
	dsn := "root:lmh123@tcp(127.0.0.1:3306)/chess?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
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
