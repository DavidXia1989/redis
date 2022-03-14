package redis_test

import (
    "github.com/DavidXia1989/redis"
    "fmt"
    "path/filepath"
    "testing"
    "time"
)
var (
    redisServerConf, _ = filepath.Abs(filepath.Join("testdata", "redis.yaml"))
    redisClient *redis.Client
)

func TestRedis(t *testing.T){
    redisClient,err := redis.InitClient(redisServerConf)
    if err!=nil {
        fmt.Println(err.Error())
    }

    // 使用默认连接
    redisClient.Set("foo", "bar from 'default' connection", 5 * time.Minute)
    fmt.Println(redisClient.Get("foo").Val())

    // 使用指定连接
    testClinet := redis.GetRedisClient("test")
    testClinet.Set("foo1", "bar from 'test' connection", 5 * time.Minute)
    fmt.Println(testClinet.Get("foo1").Val())
}