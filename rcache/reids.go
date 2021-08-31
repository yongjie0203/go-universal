package rcache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ClusterClient *redis.ClusterClient

func SetupRedis() {
	//连接服务器
	ClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"redis-node1:6379",
			"redis-node2:6379",
			"redis-node3:6379",
			"redis-node4:6379",
			"redis-node5:6379",
			"redis-node6:6379",
		},
	})

	result, error := ClusterClient.Ping(ClusterClient.Context()).Result()
	println(result, error)
	if error != nil {
		fmt.Println("redis cluster:", result)
	} else {
		fmt.Println(error)
	}
	go func() {
		for i := 0; i < 10000; i++ {
			go fmt.Println(NextId("test3"))
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			go fmt.Println(NextId("test3"))
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			go fmt.Println(NextId("test3"))
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			go fmt.Println(NextId("test3"))
		}
	}()

	for i := 0; i < 10000; i++ {
		go fmt.Println(NextId("test3"))
	}

}

func NextId(key string) uint64 {
	i := ClusterClient.Exists(ClusterClient.Context(), key)
	exists, _ := i.Result()
	if exists > 0 {
		intCmd := ClusterClient.IncrBy(ClusterClient.Context(), key, 1)
		id, _ := intCmd.Uint64()
		return id
	} else {
		ClusterClient.Set(ClusterClient.Context(), key, 1, time.Hour*24*365)
		return 1
	}
}
