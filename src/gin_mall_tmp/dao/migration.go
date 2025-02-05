/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 14:56:14
 */
package dao

import (
	"fmt"
	"mall/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Admin{},
		&model.Carousel{},
		&model.Cart{},
		&model.Category{},
		&model.Favorite{},
		&model.Notice{},
		&model.Order{},
		&model.Product{},
		&model.ProductImg{},
	)
	if err != nil {
		fmt.Println("err ", err)
	}
	return
}
