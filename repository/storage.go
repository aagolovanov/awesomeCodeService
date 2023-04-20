package repository

import (
	"context"
	"time"
)

type Storage interface {
	SetData(ctx context.Context, key string, fields map[string]string) error
	GetAllData(ctx context.Context, key string) (map[string]string, error)
	CheckExist(ctx context.Context, key string) bool
	Delete(ctx context.Context, key string) error
	SetExpire(ctx context.Context, key string, duration time.Duration) error
	Increment(ctx context.Context, key, field string) error
}
