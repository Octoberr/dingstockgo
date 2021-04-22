package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()


func ExampleClient(rdbconf *RdbConf) string {
	fmt.Printf("Redis server address: %s, port :%s, db: %d\n", rdbconf.Address, rdbconf.Port, rdbconf.DB)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "47.108.176.230:6379",
		Password: "swm123",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err ==redis.Nil {
		fmt.Println("key2 does not exist")
	}else if err != nil {
		panic(err)
	}else {
		fmt.Println("key2", val2)
	}
	return "Test OK"
}


