package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type KeyDB struct {
	rdb *redis.Client
}

var _ Storage = (*KeyDB)(nil) // for ide impl

// SetData adds data to db
func (k *KeyDB) SetData(ctx context.Context, key string, fields map[string]string) error {
	return k.rdb.HSet(ctx, key, fields).Err()
}

// GetAllData fetches data from db
func (k *KeyDB) GetAllData(ctx context.Context, key string) (map[string]string, error) {
	result := k.rdb.HGetAll(ctx, key)
	if result.Err() != nil {
		return nil, result.Err()
	} else {
		return result.Val(), nil
	}
}

func (k *KeyDB) CheckExist(ctx context.Context, key string) bool {
	result := k.rdb.Exists(ctx, key)
	return result.Err() == nil && result.Val() != 0
}

func (k *KeyDB) Delete(ctx context.Context, key string) error {
	return k.rdb.Del(ctx, key).Err() // maybe HDel
}

func (k *KeyDB) SetExpire(ctx context.Context, key string, duration time.Duration) error {
	return k.rdb.Expire(ctx, key, duration).Err() // todo return app-specific errors
}

func (k *KeyDB) Increment(ctx context.Context, key, field string) error {
	return k.rdb.HIncrBy(ctx, key, field, 1).Err() // todo return app-specific errors
}
