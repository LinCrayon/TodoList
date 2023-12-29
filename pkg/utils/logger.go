package utils

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger

func InitLog() {
	if LogrusObj != nil {
		src, _ := setOutputFile() //输出源
		// 设置输出
		LogrusObj.Out = src
		return
	}

	// 日志对象的实例化
	logger := logrus.New()
	src, _ := setOutputFile()
	// 设置输出
	logger.Out = src
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/" //日志输出的文件夹
	}
	_, err := os.Stat(logFilePath) //检查文件或目录是否存在
	if os.IsNotExist(err) {        //检查给定的错误是否表示文件或目录不存在
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	// 写入文件  os.ModeAppend 写入时将数据追加到文件的末尾。这个模式通常与追加模式一起使用,防止覆盖
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
