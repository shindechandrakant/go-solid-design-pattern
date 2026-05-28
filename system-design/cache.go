package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/singleflight"
)

var ErrorNotFound = errors.New("not found")

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, bool, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

type DB interface {
	Load(ctx context.Context, key string) ([]byte, error)
	Save(ctx context.Context, value []byte, key string) error
}

type CacheAside struct {
	cache Cache
	db    DB
	ttl   time.Duration
	group singleflight.Group
}

var negativeCacheMarker = []byte("\x00__NOT_FOUND__\x00")

func NewCacheAside(cache Cache, db DB, ttl time.Duration) *CacheAside {
	return &CacheAside{
		cache: cache,
		db:    db,
		ttl:   ttl,
	}
}

func (c *CacheAside) jitteredTTL() time.Duration {
	jitter := time.Duration(rand.Int63n(int64(c.ttl)) / 5)
	return c.ttl - (c.ttl / 10) + jitter
}

func (c *CacheAside) Get(ctx context.Context, key string) ([]byte, error) {

	if val, ok, err := c.cache.Get(ctx, key); ok && err == nil {
		if string(val) == string(negativeCacheMarker) {
			return nil, ErrorNotFound
		}
		return val, nil
	}

	val, err, _ := c.group.Do(key, func() (interface{}, error) {

		data, dbErr := c.db.Load(ctx, key)
		if errors.Is(dbErr, ErrorNotFound) {
			_ = c.cache.Set(ctx, key, negativeCacheMarker, 30*time.Second)
			return nil, ErrorNotFound
		}
		if dbErr != nil {
			return nil, dbErr
		}

		_ = c.cache.Set(ctx, key, data, c.jitteredTTL()+c.ttl)
		return data, nil
	})

	if err != nil {
		return nil, err
	}
	return val.([]byte), nil

}

func (c *CacheAside) Set(ctx context.Context, key string, value []byte) error {

	if err := c.db.Save(ctx, value, key); err != nil {
		return err
	}

	if err := c.cache.Delete(ctx, key); err != nil {
		return fmt.Errorf("db written but cache invalidation failed: %w", err)
	}
	return nil
}

func main() {

	n, err := fmt.Println("Hello")
	fmt.Println(n, err)
}
