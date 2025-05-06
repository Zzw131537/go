/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 20:25:38
 */
package model

type BasePage struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}
