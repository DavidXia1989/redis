##Redis Client Manger

解决多个连接，配置管理问题

## Installation

Env：

golang >= 1.13.0

Install:

```
// 安装
go get github.com/DavidXia1989/redis
```

## Quickstart

```go
func main(){
    // 读取配置文件 并 返回默认连接
    redisServerConf, _ := filepath.Abs(filepath.Join("conf", "redis.yaml"))
    redisClient,_:=redis.InitClient(redisServerConf)


    // 使用默认连接
    redisClient.Set("foo", "bar from 'default' connection", 5 * time.Minute)
    fmt.Println(redisClient.Get("foo").Val())

    // 使用指定连接
    testClinet := redis.GetRedisClient("test")
    testClinet.Set("foo1", "bar from 'test' connection", 5 * time.Minute)
    fmt.Println(testClinet.Get("foo1").Val())
}
```

## Howto

基于[https://github.com/go-redis/redis](https://github.com/go-redis/redis)开发，所有方法请参考官方文档

## Config

使用ymal格式配置

```yaml
-
  name: default
  host: 127.0.0.1
  port: 6379
  password:
  db:
- name: test
  host: 127.0.0.1
  port: 6379
  password:
  db:
```

