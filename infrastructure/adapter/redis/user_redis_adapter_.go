package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	c "rpolnx.com.br/golang-hex/application/config"
	ce "rpolnx.com.br/golang-hex/application/error"
	u "rpolnx.com.br/golang-hex/domain/model/user"
	"rpolnx.com.br/golang-hex/domain/ports/out"
)

const (
	PATTERN = "USER:%s"
)

type redisRepository struct {
	client  *redis.Client
	cache   *cache.Cache
	timeout time.Duration
	ttl     time.Duration
}

func InitializeRepo(config c.Redis) (u out.RedisUserRepository, err error) {
	repo, err := newRepository(config)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func newRepository(config c.Redis) (out.RedisUserRepository, error) {
	repo := &redisRepository{
		timeout: time.Duration(config.Timeout) * time.Second,
		ttl:     time.Duration(config.Ttl) * time.Minute,
	}

	client, err := newClient(config)
	if err != nil {
		return nil, errors.Wrap(err, "repository.redis.newRepository")
	}
	repo.client = client

	c := cache.New(&cache.Options{
		Redis: client,
	})

	repo.cache = c

	return repo, nil
}

func newClient(config c.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         config.Host,
		Password:     config.Password,
		Username:     config.Username,
		DB:           config.Db,
		PoolSize:     config.PoolSize,
		PoolTimeout:  time.Duration(config.Timeout) * time.Second,
		ReadTimeout:  time.Duration(config.Timeout) * time.Second,
		WriteTimeout: time.Duration(config.Timeout) * time.Second,
		IdleTimeout:  time.Duration(config.IdleTimeout) * time.Second,
	})

	res := rdb.Ping(context.Background())

	if res == nil {
		return nil, errors.New("Error connecting to redis")
	}

	return rdb, nil
}

func (r *redisRepository) Get(name string) (*u.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	user := &u.User{}

	err := r.cache.Get(ctx, fmt.Sprintf(PATTERN, name), &user)
	if err != nil {
		if err == cache.ErrCacheMiss {
			return nil, ce.ErrCacheMiss
		}
		return nil, errors.Wrap(err, "repository.redis.Get")
	}
	return user, nil
}

func (r *redisRepository) Set(user *u.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	err := r.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf(PATTERN, user.Name),
		Value: user,
		TTL:   r.ttl,
	})

	if err != nil {
		return errors.Wrap(err, "repository.redis.Set")
	}

	return nil
}
