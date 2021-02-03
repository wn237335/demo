package main

import (
	"context"
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

	err := rdb.Set(ctx, "ceping_othersdk", "fdgdgf", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	_ = rdb.Close()
}
