package redishelper

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func RedisGroupRead() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "8.209.247.215:6380", // Redis 服务器地址
		Password: "4MuUXAVfp",          // Redis 服务器密码
		DB:       2,                    // 使用默认数据库
	})

	// 关闭连接
	defer client.Close()

	// 消费者组名称
	groupName := "my_consumer_group"
	// 流名称
	streamName := "mystream"
	// 消费者名称
	consumerName := "my_consumer"

	// 创建消费者组
	if err := client.XGroupCreateMkStream(context.Background(), streamName, groupName, "$").Err(); err != nil {
		fmt.Println("Error creating consumer group:", err)
		return
	}

	// 从流中读取消息
	for {
		// 读取消息
		msgs, err := client.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Group:    groupName,
			Consumer: consumerName,
			Streams:  []string{streamName, ">"},
			Count:    10,   // 每次读取的消息数量
			Block:    0,    // 阻塞时间，设置为0表示不阻塞
			NoAck:    true, // 不需要手动确认消息
		}).Result()

		if err != nil {
			fmt.Println("Error reading messages:", err)
			return
		}

		// 处理读取到的消息
		for _, stream := range msgs {
			streamName := stream.Stream
			fmt.Printf("Stream: %s\n", streamName)
			for _, msg := range stream.Messages {
				messageID := msg.ID
				messageValues := msg.Values
				fmt.Printf("Message ID: %s, Values: %v\n", messageID, messageValues)
			}
		}
	}
}
