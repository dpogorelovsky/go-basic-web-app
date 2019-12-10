package redis

import (
	"fmt"
	"strings"
	"time"

	"github.com/dpogorelovsky/go-basic-web-app/app/cache/acache"
	"github.com/go-redis/redis"
)

// Client - redis client instance
type Client struct {
	driver *redis.Client
	acache.AbstractCache
}

// NewRedisClient - redis client constructor
func NewRedisClient(host, port, password string, db int) (*Client, error) {
	addr := fmt.Sprintf("%s:%s", host, port)
	rds := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := rds.Ping().Result()
	if err != nil {
		return nil, err
	}

	cln := &Client{driver: rds}

	return cln, nil
}

// Get - get value by key
func (c *Client) Get(key string) (string, error) {
	res, err := c.driver.Get(key).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

// Set - value by key
func (c *Client) Set(key, val string, ttl time.Duration) error {
	expires := time.Second * ttl
	_, err := c.driver.Set(key, val, expires).Result()

	return err
}

// GenerateKey - cache key generator
func (c *Client) GenerateKey(args ...string) string {
	cacheDelim := ":"
	argsStr := strings.Join(args, cacheDelim)

	return fmt.Sprintf("%s%s%s", "cache", cacheDelim, argsStr)
}
