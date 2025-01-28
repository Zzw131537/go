/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-28 16:57:25
 */
package ws

type Trainer struct {
	Content   string `bson:"content"`   // 内容
	StartTime int64  `bson:"startTime"` // 创建时间
	EndTime   int64  `bson:"endTime"`   // 结束时间
	Read      uint   `bason:"read"`     // 已读
}

type Result struct {
	StartTime int64
	Msg       string
	Content   interface{}
	From      string
}
