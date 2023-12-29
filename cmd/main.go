package main

import (
	"todo_list_v2.01/config"
	"todo_list_v2.01/pkg/utils"
	"todo_list_v2.01/repository/cache"
	"todo_list_v2.01/repository/db/dao"
	"todo_list_v2.01/router"
)

func main() {
	loading()
	r := router.NewRouter()
	_ = r.Run(config.Config.System.HttpPort)
}
func loading() {
	config.InitConfig()
	dao.MySQLInit()
	utils.InitLog()
	cache.RedisInit()
}
