/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-20 17:57:07
 */
package model

func Migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{})
}
