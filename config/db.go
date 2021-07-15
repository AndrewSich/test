package config

import (
	"fmt"
	_ "github.com/jinzhu/dialects/mysql"
	"github.com/jinzhu/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Lupakatasandi1711-@tcp(34.87.101.85:3306)/flome?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("[FLOME] (Init) ", err)
	} else {
		fmt.Println("[FLOME] (Init) Success connect to database")
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

func TestDBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Lupakatasandi1711-@tcp(34.87.101.85:3306)/flome_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("[FLOME] (Test Init) ", err)
	} else {
		fmt.Println("[FLOME] (Test Init) Success connect to Database")
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
