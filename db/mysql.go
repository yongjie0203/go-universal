package db

import (
	"github.com/yongjie0203/go-universal/rcache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var Coon *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:root@(127.0.0.1:3310)/trade?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})

	if err != nil {
		log.Print(err.Error())
	}

	sqldb, e := db.DB()
	if e != nil {
		log.Print(e.Error())
	}

	e = sqldb.Ping()
	if e != nil {
		panic(e)
	}
	sqldb.SetMaxIdleConns(100)
	sqldb.SetMaxOpenConns(20)
	Coon = db
	return db
}

var updateDatabaseHandlerMap = make(map[string]interface{})

func TableUpdateRegister(name string, o interface{}) {
	updateDatabaseHandlerMap[name] = o
}
func DBUpdate() {
	for name, obj := range updateDatabaseHandlerMap {
		log.Printf("update table %s", name)
		Coon.AutoMigrate(obj)
	}

}

var id int64

// NextId 可更换为Redis实现
func NextId(tableName string) int64 {
	return int64(rcache.NextId(tableName))
}
