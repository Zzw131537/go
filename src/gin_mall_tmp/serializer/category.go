/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 21:43:40
 */
package serializer

import "mall/model"

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(iten *model.Category) Category {
	return Category{
		Id:           iten.ID,
		CategoryName: iten.CategoryName,
		CreateAt:     iten.CreatedAt.Unix(),
	}
}

func BuildCategorys(items []*model.Category) (categorys []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categorys = append(categorys, category)
	}
	return
}
