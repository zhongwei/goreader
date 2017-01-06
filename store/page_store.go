package store

import (
    "log"

    "github.com/garyburd/redigo/redis"
)

func Save(url, content string) {
    c, err := redis.Dial("tcp", "redis:6379")
    if err != nil {
        panic(err)
    }
    defer c.Close()
    c.Do("SET", url, content)
    readContent, err := redis.String(c.Do("GET", url))
    log.Println("url=>content : \n" + url + "\n" + readContent)
}
