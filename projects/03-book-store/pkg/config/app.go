package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "durga:Durga@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Warning: Could not connect to MySQL -", err.Error())
		fmt.Println("Database features will be unavailable until MySQL is running")
		return
	}
	db = d
}

func GetDB() *gorm.DB {
	if db == nil {
		fmt.Println("Database not connected. Please start MySQL and restart the application.")
	}
	return db
}
