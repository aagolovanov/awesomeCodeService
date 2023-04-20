package repository

import (
	"context"
	"time"
)

type KeyDB struct {
}

var _ Storage = (*KeyDB)(nil) // for ide impl

func (k *KeyDB) SetData(ctx context.Context, key string, fields map[string]string) error {
	//TODO implement me
	panic("implement me")
}

func (k *KeyDB) GetData(ctx context.Context, key string) {
	//TODO implement me
	panic("implement me")
}

func (k *KeyDB) CheckExist(ctx context.Context, key string) bool {
	//TODO implement me
	panic("implement me")
}

func (k *KeyDB) Delete(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}

func (k *KeyDB) SetExpire(ctx context.Context, key string, duration time.Duration) error {
	//TODO implement me
	panic("implement me")
}
