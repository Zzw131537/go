/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-15 16:53:59
 */
package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// user 错误
	ErrorExistUser             = 30001
	ErrorFailEncryption        = 30002
	ErrorExistUserNotFound     = 30003
	ErrorNotCompare            = 30004
	ErrorAuthToken             = 30005
	ErrorAuthCheckTokenTimeout = 30006
	ErrorUpLoadFail            = 30007
	ErrorSendEmail             = 30008
	// produce 模块错误 4xxxxx
	ErrorProductImgUpload = 40001

	// 收藏夹错误
	ErrorFavoriteExist = 50001
)
