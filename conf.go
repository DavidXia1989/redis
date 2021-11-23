package redis

import "time"

type RedisConf struct {
    Host         string        `json:"host"`
    Name         string        `json:"name"`
    Port         string        `json:"port"`
    Password     string        `json:"password"`
    DB           int           `json:"db"`
    DialTimeout  time.Duration `json:"dial_timeout"`
    ReadTimeout  time.Duration `json:"read_timeout"`
    WriteTimeout time.Duration `json:"write_timeout"`
    PoolSize     int           `json:"pool_size"`
    PoolTimeout  time.Duration `json:"pool_timeout"`
}

func NewRedisConf() *RedisConf {
    return &RedisConf{
        Host:         "",
        Name:         "",
        Port:         "",
        Password:     "",
        DB:           0,
        DialTimeout:  10 * time.Second,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        PoolSize:     10,
        PoolTimeout:  30 * time.Second,
    }
}

func (c *RedisConf) buildDefault() {
    if c.DialTimeout == 0 {
        c.DialTimeout = 10 * time.Second
    }
    if c.ReadTimeout == 0 {
        c.ReadTimeout = 30 * time.Second
    }
    if c.WriteTimeout == 0 {
        c.WriteTimeout = 30 * time.Second
    }
    if c.PoolSize == 0 {
        c.PoolSize = 10
    }
    if c.PoolTimeout == 0 {
        c.PoolTimeout = 30 * time.Second
    }
}
