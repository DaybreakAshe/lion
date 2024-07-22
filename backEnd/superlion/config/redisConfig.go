package config

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const (
	Redis_host     = "caching-1bbaeb26-dravenxue-82bd.h.aivencloud.com"
	Redis_port     = "15278"
	Redis_UserName = "default"
	Redis_pwd      = "AVNS_cDEr2rLy4i2d2q1gX3l"
	Redis_database = 0

	//Redis_host     = "yanjl.eu.org"
	//Redis_port     = "16379"
	//Redis_pwd      = ""
	//Redis_database = 2
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
		PoolSize:     5,
		PoolTimeout:  30 * time.Second,
		TLSConfig: &tls.Config{
			// 根据需要配置 TLS 选项
			InsecureSkipVerify: true, // 如果你没有 CA 证书，可以临时跳过证书验证
		},
	})

	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client = redisDB
		redisClient = rdh
	})
	fmt.Println("redis init over,everything is OK")
	return redisDB
}
