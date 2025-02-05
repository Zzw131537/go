/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 15:25:32
 */
package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
