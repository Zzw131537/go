/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-06 15:08:31
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
	Token string      `json:"tokens"`
}
