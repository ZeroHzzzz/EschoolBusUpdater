package redis

import (
	"github.com/go-redis/redis"
)

type redisConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

var RedisClient *redis.Client
var RedisInfo redisConfig

func init() {
	info := getConfig()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     info.Host + ":" + info.Port,
		Password: info.Password,
		DB:       info.DB,
	})
	RedisInfo = info

}

func TestConnection() (bool,error) {
	_, err := RedisClient.Ping().Result()
	return err==nil, err
} 