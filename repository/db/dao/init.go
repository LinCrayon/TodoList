package dao

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
	"todo_list_v2.01/config"
)

var _db *gorm.DB

func MySQLInit() {
	mConfig := config.Config.MySql["default"]
	conn := strings.Join([]string{mConfig.User, ":", mConfig.Password,
		"@tcp(", mConfig.Host, ":", mConfig.Port, ")/", mConfig.DbName,
		"?charset=utf8mb4&parseTime=true"}, "")
	fmt.Println(conn + "==================================")
	//日志记录器
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,  // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger, // 打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表明不加s
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	migration()
}

// NewDBClient 指定上下文的数据库连接实例
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx) //基于传入的 context对象创建了一个新的数据库连接实例
} //可以确保在数据库操作中使用相同的上下文，以便在需要时进行取消或超时处理。
