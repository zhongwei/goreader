package store

import (
    "log"

    "github.com/garyburd/redigo/redis"
    "github.com/spf13/viper"
)

func Save(url, content string) {
    viper.SetDefault("redis.server", "localhost")
    viper.SetDefault("redis.port", "6379")
    viper.SetDefault("redis.protocol", "tcp")

    err := viper.ReadInConfig()
    server := viper.GetString("redis.server")
    port := viper.GetString("redis.port")
    protocol := viper.GetString("redis.protocol")
    
    log.Println("redis server: " + protocol + "://" + server + ":" + port)

    c, err := redis.Dial(protocol, server + ":" + port)
    if err != nil {
        panic(err)
    }
    defer c.Close()
    c.Do("SET", url, content)
}
