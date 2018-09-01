package models

import "fmt"

//RedisTest 使用demo
func RedisTest() {
	err := redisClient.HSet("data", "key", "value1111111").Err()
	if err != nil {
		fmt.Println("hset")
		fmt.Println(err)
	} else {
		rest, err := redisClient.HGet("data", "key").Result()
		if err != nil {
			fmt.Println("hget")
			fmt.Println(err)
		}
		fmt.Println(rest)
	}
}
