/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-09 17:02:56
 */
package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  item,
			Total: total,
		},
		Msg: "ok",
	}
}
