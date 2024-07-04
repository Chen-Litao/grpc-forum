package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
)

var Ctx = context.Background()

var RdbTest *redis.Client

// UserFollowings 根据用户id找到他关注的人
var UserFollowings *redis.Client

// UserFollowers 根据用户id找到他的粉丝
var UserFollowers *redis.Client

func InitRedis() {
	addr := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	RdbTest = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	UserFollowings = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       1,
	})
	UserFollowers = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       2,
	})
	_, err := UserFollowings.Ping(context.Background()).Result()
	if err != nil {
		log.Panicf("连接 redis 错误，错误信息: %v", err)
	} else {
		log.Println("Redis 连接成功！")
	}
}
