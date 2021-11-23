package redis

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var redisClients = make(map[string]*Client)

type Client struct {
	*redis.Client
}

func InitClient(confFile string) (*Client, error) {
	var conf []RedisConf
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		// 配置文件未找到
		return nil, errors.New("redis 配置文件读取失败：" + err.Error())
	}

	{
		err := yaml.Unmarshal(yamlFile, &conf)
		if err != nil {
			// 配置文件解析失败
			return nil, errors.New("redis 配置文件解析失败：" + err.Error())
		}
	}

	var c *Client
	for k := range conf {
		conf[k].buildDefault()
		var client *redis.Client
		client = redis.NewClient(&redis.Options{
			Addr:         conf[k].Host + ":" + conf[k].Port,
			Password:     conf[k].Password,
			DB:           conf[k].DB,
			DialTimeout:  conf[k].DialTimeout,
			ReadTimeout:  conf[k].ReadTimeout,
			WriteTimeout: conf[k].WriteTimeout,
			PoolSize:     conf[k].PoolSize,
			PoolTimeout:  conf[k].PoolTimeout,
		})
		redisClients[conf[k].Name] = &Client{client}
		if err := redisClients[conf[k].Name].Ping().Err(); err != nil {
			// pinc
			return nil, errors.New("redis 连接失败，host：" + conf[k].Host + "，err：" + err.Error())
		}
		if conf[k].Name == "default" {
			c = redisClients[conf[k].Name]
		}
	}
	return c, nil
}
func NewClients(conf []RedisConf) error {
	for k := range conf {
		var err error
		redisClients[conf[k].Name], err = NewClient(conf[k])
		if err != nil && redisClients[conf[k].Name] == nil {
			return errors.New("RedisGroup is error：" + err.Error())
		}
	}
	return nil
}
func NewClient(c RedisConf) (*Client, error) {
	if redisClients[c.Name] != nil {
		return redisClients[c.Name], errors.New("连接名已存在")
	}
	c.buildDefault()
	redisClients[c.Name] = &Client{redis.NewClient(&redis.Options{
		Addr:         c.Host + ":" + c.Port,
		Password:     c.Password,
		DB:           c.DB,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		PoolSize:     c.PoolSize,
		PoolTimeout:  c.PoolTimeout,
	})}
	if err := redisClients[c.Name].Ping().Err(); err != nil {
		// pinc
		return nil, errors.New("redis 连接失败，host：" + c.Host + "，err：" + err.Error())
	}
	return redisClients[c.Name], nil
}

func GetRedisClient(name string) *Client {
	if redisClients[name] != nil {
		return redisClients[name]
	}
	return nil
}
