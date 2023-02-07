package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type IRedis interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, dest interface{}) error
	GetString(ctx context.Context, key string) (res string, err error)
	Incr(ctx context.Context, key string) (int64, error)
	Del(ctx context.Context, key string) error
	DelBulk(ctx context.Context, prefix string) error
	SetIfNotExist(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	HashSet(ctx context.Context, hashName string, value ...interface{}) (int64, error)
	HashGet(ctx context.Context, hashName, field string) (string, error)
	HashDel(ctx context.Context, hashName string, fields ...string) (int64, error)
	Lock(ctx context.Context, key, identifier string, expiration time.Duration) bool
	Unlock(ctx context.Context, key, identifier string)
}

type Client struct {
	client *redis.Client
}

func NewClient(c *redis.Client) *Client {
	return &Client{
		client: c,
	}
}

func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, b, expiration).Err()
}

func (c *Client) Get(ctx context.Context, key string, dest interface{}) error {
	var data []byte
	err := c.client.Get(ctx, key).Scan(&data)
	if err != nil {
		if err == redis.Nil {
			return nil
		}
	}

	return json.Unmarshal(data, dest)
}

func (c *Client) GetString(ctx context.Context, key string) (res string, err error) {
	res, err = c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return res, nil
		}
	}

	return res, err
}

func (c *Client) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

func (c *Client) Del(ctx context.Context, key string) error {
	_, err := c.client.Del(ctx, key).Result()
	return err
}

func (c *Client) DelBulk(ctx context.Context, prefix string) error {
	iter := c.client.Scan(ctx, 0, prefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		_, err := c.client.Del(ctx, iter.Val()).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

// SetIfNotExist will return a boolean value explaining whether or not the key already exist;
// and the set operation is not executed, or an error
func (c *Client) SetIfNotExist(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.client.SetNX(ctx, key, value, expiration).Result()
}

func (c *Client) HashSet(ctx context.Context, hashName string, value ...interface{}) (int64, error) {
	return c.client.HSet(ctx, hashName, value).Result()
}

func (c *Client) HashGet(ctx context.Context, hashName, field string) (string, error) {

	res, err := c.client.HGet(ctx, hashName, field).Result()
	if err != nil {
		if err == redis.Nil {
			return res, nil
		}
		return res, err
	}

	return res, nil
}

func (c *Client) HashDel(ctx context.Context, hashName string, fields ...string) (int64, error) {
	return c.client.HDel(ctx, hashName, fields...).Result()
}

func (c *Client) Lock(ctx context.Context, key, identifier string, expiration time.Duration) bool {
	timeout := time.Now().Add(expiration)
	for time.Now().Before(timeout) {
		if c.client.SetNX(ctx, key, identifier, expiration).Val() {
			return true
		}
		time.Sleep(1 * time.Millisecond)
	}

	return false
}

func (c *Client) Unlock(ctx context.Context, key, identifier string) {
	res, _ := c.client.Get(ctx, key).Result()
	if res == identifier {
		c.client.Del(ctx, key)
	}
}
