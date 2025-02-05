/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-04 18:04:54
 */
package model

type BasePage struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}
