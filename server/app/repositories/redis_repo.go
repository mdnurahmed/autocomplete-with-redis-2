package repositories

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type IRedisRepository interface {
	Search(key string, length int64) (result []string, err error)
	Insert(word, key string) (err error)
	Delete() (err error)
}

type RedisRepository struct {
	rdb *redis.Client
}

func NewInstanceOfRedisRepository(
	address,
	password string,
	DB int) RedisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       DB,
	})
	return RedisRepository{
		rdb: rdb,
	}
}

func (a *RedisRepository) Insert(word, key string) error {
	length := a.rdb.ZCard(key).Val()
	pipe := a.rdb.Pipeline()
	if length >= 300 {
		res, err := a.rdb.ZPopMin(key).Result()
		if err != nil {
			return err
		}
		if len(res) == 0 {
			return fmt.Errorf("element with lowest score could not be retrieved")
		}
		pipe.ZAdd(key, redis.Z{
			Score:  res[0].Score + 1,
			Member: word,
		})
		pipe.Expire(key, time.Hour)
	} else {
		pipe.ZIncrBy(key, 1.0, word)
		pipe.Expire(key, time.Hour)
	}

	_, err := pipe.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (a *RedisRepository) Search(key string, searchLength int64) ([]string, error) {
	result := a.rdb.ZRevRange(key, 0, searchLength)
	searchResult, err := result.Result()
	return searchResult, err
}

func (a *RedisRepository) Delete() error {
	a.rdb.FlushDBAsync()
	return nil
}
