/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-16 21:28:11
 */
package cache

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	RedisClient *redis.Client

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./congig/config.ini") // 加载配置信息文件
	if err != nil {
		fmt.Println("ini Load failed", err)

	}
	LoadRedis(file) // 读取文件内容
	Redis()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64) // string to 64
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
