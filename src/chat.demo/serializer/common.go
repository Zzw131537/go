/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-20 17:18:43
 */
package serializer

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
