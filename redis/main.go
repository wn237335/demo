package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"redis/utils"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     utils.Sspredisserverandport(),
		Password: utils.Sspredispwd(), // no password set
		DB:       0,                   // use default DB
	})
	var ctx = context.Background()
	items := make(map[string]interface{})
	items["username"] = "root"
	items["password"] = "123456"
	ss, _ := json.Marshal(items)

	//存入一个值
	err := rdb.Set(ctx, "test", ss, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	//1查询一个key
	val, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("test", val)

	//2查询一个key
	val2, err := rdb.Get(ctx, "test").Result()
	if err == redis.Nil {
		fmt.Println("test does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("test", val2)
	}

	//关闭链接
	_ = rdb.Close()
}
