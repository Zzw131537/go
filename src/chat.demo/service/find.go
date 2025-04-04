/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-28 17:52:18
 */
package service

import (
	"chat/conf"
	"chat/model/ws"
	"context"
	"time"
)

func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	collection := conf.MongoDBClient.Database(database).Collection(id)
	comment := ws.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(), comment)
	return
}
