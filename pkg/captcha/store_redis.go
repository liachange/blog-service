package captcha

import (
	"blog-service/pkg/app"
	"blog-service/pkg/config"
	"blog-service/pkg/redis"
	"errors"
	"time"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (s *RedisStore) Set(key, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}
	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	value := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return value
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	value := s.Get(key, clear)
	return answer == value
}
