package verifycode

import (
	"blog-service/pkg/app"
	"blog-service/pkg/config"
	"blog-service/pkg/redis"
	"time"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 实现 verifycode.Store interface 的 Set 方法
func (s *RedisStore) Set(key, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}
	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

// Get 实现 verifycode.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	value := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return value
}

// Verify 实现 verifycode.Store interface 的 Verify 方法
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	value := s.Get(id, clear)
	return value == answer
}
