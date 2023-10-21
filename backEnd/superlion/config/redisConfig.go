package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const (
	//Redis_host     = "redis-14382.c81.us-east-1-2.ec2.cloud.redislabs.com"
	//Redis_port     = "14382"
	//Redis_pwd      = "ROOT123"
	//Redis_database = 0

	Redis_host     = "148.100.77.194"
	Redis_port     = "39379"
	Redis_pwd      = ""
	Redis_database = 6
)

type RedisHelper struct {
	*redis.Client
}

var redisClient *RedisHelper

var redisOnce sync.Once

func GetRedisHelper() *RedisHelper {
	return redisClient
}

func NewRedisHelper() *redis.Client {
	redisDB := redis.NewClient(&redis.Options{
		Addr:         Redis_host + ":" + Redis_port,
		Password:     Redis_pwd,
		DB:           Redis_database,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client = redisDB
		redisClient = rdh
	})
	fmt.Println("redis init over,everything is OK")
	return redisDB
}
