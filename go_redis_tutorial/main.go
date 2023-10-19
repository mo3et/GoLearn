package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// 连接Redis

// Background返回一个非空的Context。它永远不好被取消，没有值，也没有期限。
// 它通常在main函数，初始化和测试时使用，并用作传入请求的顶级Context (Parent Context)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		// Addr:     "172.16.147.128:26379",
		Addr:     "127.0.0.1:26379",
		Password: "redis",
		DB:       0,
		PoolSize: 100,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Connection Redis error: %v\n", err)
	}
	for i := 0; i < 10; i++ {
		x := strconv.Itoa(i)
		err = rdb.Set(ctx, x, x+" hello", 0).Err()
		if err != nil {
			panic(err)
		}

	}

	for i := 0; i < 8; i++ {
		x2 := strconv.Itoa(i)
		val, err := rdb.Get(ctx, x2).Result()
		if err != nil {
			fmt.Printf("Dont found key: %v\n", err)
		}
		fmt.Println(val)
	}
	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	fmt.Printf("Dont found key: %v\n", err)
	// }
	fmt.Println("Connection Success", pong)
}

type Config struct {
	// Network "tcp"
	Network string
	// Addr "127.0.0.1:6379"
	Addr string
	// Password string .If no password then no 'AUTH'. Default ""
	Password string
	// If Database is empty "" then no 'SELECT'. Default ""
	Database string
	// MaxIdle 0 no limit
	MaxIdle int
	// MaxActive 0 no limit
	MaxActive int
	// IdleTimeout  time.Duration(5) * time.Minute
	IdleTimeout time.Duration
	// Prefix "myprefix-for-this-website". Default ""
	Prefix string
}
