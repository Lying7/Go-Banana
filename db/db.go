package db

import (
	"Go-Banana/config"
	"Go-Banana/glog"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var myDb *DB

func init() {
	//dsn := "admin:lpqq456Q!@tcp(106.53.153.76:3306)/GoBanana?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	myDb, err = connectDB(config.Config().DB.User,
		config.Config().DB.Passwd,
		config.Config().DB.Addr,
		config.Config().DB.Port,
		config.Config().DB.DBName)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}

	glog.Debug("Connect db success")
}

func connectDB(user string, passwd string, svrAddr string, svrPort int, dbName string) (db *DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, passwd, svrAddr, svrPort, dbName)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
