package config

import (
	"github.com/spf13/viper"
	"os"
)

var Config *Conf

type Conf struct {
	System *System           `yaml:"system"`
	MySql  map[string]*MySql `yaml:"mysql"`
	Redis  *Redis            `yaml:"redis"`
}

type MySql struct {
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbName"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Redis struct {
	RedisHost string `yaml:"redisHost"`
	RedisPort string `yaml:"redisPort"`
	//RedisPassword string `yaml:"redisPassword"`
	RedisDbName  int    `yaml:"redisDbName"`
	RedisNetwork string `yaml:"redisNetwork"`
}

type System struct {
	AppEnv   string `yaml:"appEnv"`
	Domain   string `yaml:"domain"`
	Version  string `yaml:"version"`
	HttpPort string `yaml:"httpPort"`
	Host     string `yaml:"host"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/local") //文件搜索路径
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config) //将配置信息解析到指定的结构体
	if err != nil {
		panic(err)
	}
}
