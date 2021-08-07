package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "dev1:dev_test@tcp(127.0.0.1:3306)/local?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("%v", err)
	} else {
		fmt.Printf("successfully connected")
	}

	sqlDB, err := db.DB()

	for {
		if err = sqlDB.Ping(); err != nil {
			fmt.Printf("error cuk !")
		}
	}

	// db.DB().ping

}
