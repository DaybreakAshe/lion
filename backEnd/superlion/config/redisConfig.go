package config

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const (
	Redis_host     = "lion-yanjl002.h.aivencloud.com"
	Redis_port     = "24584"
	Redis_UserName = "default"
	Redis_pwd      = "AVNS_VkjQjjVxvSdhcN1gndY"
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

/*
https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
0,0,0,1,1,1,2,3,3
0 0 1 1 2 3 2 3 3
*/
func removeDuplicates(nums []int) int {

	length := len(nums)

	if length < 2 {
		return length
	}
	// 重复计数
	cnt := 1
	// 当前指针-比较基准
	head := 0
	// 移动指针
	tail := 1

	for tail < length {
		// 相等则计数+1
		if nums[head] == nums[tail] {
			cnt++
			// 计数不到2，移动基准指针，直到达到2个重复值，head+1位需要被覆盖
			if cnt <= 2 {
				// 少于2个时，直接覆盖
				nums[head+1] = nums[tail]
				head++
			}
		} else { // 不相等覆盖基准指针，重置计数
			nums[head+1] = nums[tail]
			head++
			cnt = 1
		}
		tail++
	}

	return head + 1
}
