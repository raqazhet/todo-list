package redisC

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type RedisMethod interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
}
type RedisCashe struct {
	Client *redis.Client
}

func NewCasheRedis() *RedisCashe {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("error in redisPing: %v", err)
		return nil
	}
	return &RedisCashe{
		Client: client,
	}
}

func (cl *RedisCashe) Set(key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		log.Printf("set redis err: %v", err)
		return err
	}
	if err := cl.Client.Set(key, data, ttl).Err(); err != nil {
		log.Printf("err in redis Set method: %v", err)
		return err
	}
	return nil
}

func (cl *RedisCashe) Get(key string) (interface{}, error) {
	data, err := cl.Client.Get(key).Result()
	if err != nil {
		log.Printf("err in redis Get() %v", err)
		return nil, err
	}
	var vluse interface{}
	if err := json.Unmarshal([]byte(data), &vluse); err != nil {
		return nil, err
	}
	return vluse, nil
}
