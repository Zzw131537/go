package service

import (
	"context"
	"mall/dao"
	"mall/pkg/e"
	"mall/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)

	code := e.Success

	categories, err := categoryDao.ListCategory()
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategorys(categories), uint(len(categories)))
}
