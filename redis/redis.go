package redis_cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	dataExpiration = time.Hour * 12
)

type RedisAdapter struct {
	Client *redis.Client
}

func NewRedisAdapter(client *redis.Client) *RedisAdapter {
	return &RedisAdapter{
		Client: client,
	}
}

func getKey(query, limit string) string {
	return fmt.Sprint(`SELECT * FROM Page WHERE MATCH ( CONTENT ) AGAINST (? IN NATURAL LANGUAGE MODE) LIMIT ?;`, query, limit)
}

func (r *RedisAdapter) SetSearch(query, limit string, value []byte) error {
	return r.Client.Set(
		getKey(query, limit), value, dataExpiration,
	).Err()
}

func (r *RedisAdapter) GetSearch(query, limit string) ([]byte, error) {
	return r.Client.Get(getKey(query, limit)).Bytes()
}
