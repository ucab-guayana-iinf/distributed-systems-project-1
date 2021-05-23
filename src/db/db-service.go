package main

import (
	"fmt"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "gorm.io/driver/sqlite"
)

type Count struct {
  gorm.Model
  value int
}

var counts []Count
var count_record Count

func Initialize() {
	// NOTE: checks if db exists & if value is present
	// if not initializes count in 0
	db, err := gorm.Open(
		sqlite.Open("count.db"),
		&gorm.Config{ Logger: logger.Default.LogMode(logger.Silent) })

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Count{})

	var counts []Count
	result := db.Find(&counts)

	if result.RowsAffected == 0 {
		fmt.Println("[Database]: No count records found initializing count")
		db.Create(&Count{value: 0})
	}

	// --
	result = db.First(&count_record)
	fmt.Printf("[Database]: Result %v\n", count_record.value)
	// --

	// default_count := Count{value: 0}
	// result := create()

	// fmt.Println(result)
	
	// db.First(&product, 1)

	// fmt.Println(product)
	// fmt.Println(product.value)
}

func main() {
	Initialize()
}

// func GetCount() {
// 	result := db.First(&count_record)
// 	fmt.Println("[Database]: Result %v", result)
// }

// func UpdateCount() {

// }
