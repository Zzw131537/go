/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-16 16:07:34
 */
package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserId:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Addtrss: service.Address,
	}
	err := addressDao.CreateAddress(address)
	if err != nil {
		code := e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Show(ctx context.Context, uid uint, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByaId(uint(addressId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)

	addresses, err := addressDao.ListAddressByuId(uId)
	if err != nil {
		code := e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddresses(addresses),
	}
}

func (service *AddressService) Update(ctx context.Context, uId uint, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)

	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		Name:    service.Name,
		Phone:   service.Phone,
		Addtrss: service.Address,
	}
	err := addressDao.UpdateAddressByUserId(address, uId, uint(addressId))
	if err != nil {
		code := e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)

	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	err := addressDao.DeleteAddress(uId, uint(addressId))
	if err != nil {
		util.LogrusObj.Infoln("err ", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
