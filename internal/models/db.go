package models

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() {
	username := "root"              //账号
	password := readLocalPassword() //密码
	host := "127.0.0.1"             //数据库地址，可以是Ip或者域名
	port := 3306                    //数据库端口
	Dbname := "osys"                //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
	defer initDB()
}

type config struct {
	Password string `yaml:"password"`
}

func readLocalPassword() string {
	var conf config
	result, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(result, &conf)
	if err != nil {
		panic(err)
	}
	return conf.Password
}

func initDB() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic("auto migrate err: " + err.Error())
	}
}
