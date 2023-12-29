package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"todo_list_v2.01/config"
)

var RedisClient *redis.Client

func RedisInit() {
	rConfig := config.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", rConfig.RedisHost, rConfig.RedisPort),
		//Password: rConfig.Password,
		DB: rConfig.RedisDbName,
	})

	_, err := client.Ping().Result()
	if err != nil {
		logrus.Info(err)
		panic(err)
	}
	RedisClient = client
}
