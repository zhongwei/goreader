package store

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

func Save(url, content string) {
	connection, _ := getConnection()
	defer connection.Close()
	connection.Do("SET", url, content)
	connection.Do("SET", getMD5Str(url), url)
}

func getConnection() (connection redis.Conn, err error) {
	viper.SetDefault("redis.server", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.protocol", "tcp")

	server := viper.GetString("redis.server")
	port := viper.GetString("redis.port")
	protocol := viper.GetString("redis.protocol")

	connection, err = redis.Dial(protocol, server+":"+port)
	if err != nil {
		panic(err)
	}

	log.Println("redis server: " + protocol + "://" + server + ":" + port)

	return
}

func getMD5Str(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
