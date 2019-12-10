package acache

import "time"

// AbstractCache - abstract cache provider
type AbstractCache interface {
	Get(key string) (string, error)
	Set(key, val string, ttl time.Duration) error
	GenerateKey(args ...string) string
}
