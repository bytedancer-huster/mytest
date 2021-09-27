package dao

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var userDB *gorm.DB
var redisClient *redis.Client

func init() {
	initUserDB()
	initRedis()
}

func initUserDB() {
	var err error
	userDB, err = gorm.Open("mysql", "root:gotest@(127.0.0.1:3306)/userinfo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // use default Addr
		Password: "",               // no password set
	})
}

func GetRedis() *redis.Client {
	return redisClient
}
