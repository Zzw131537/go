package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	err = db.Model(&Product{}).First(&product, 1).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(product)

	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	var product2 Product
	err = db.First(&product2, 1).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(product2)

	db.Delete(&product, 1)
	err = db.Model(&Product{}).First(&product2, 1).Error
	if err != nil {
		panic(err)
	}
}
