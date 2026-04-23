package rate_limiter

import (
	"sync"
	"time"
)

type RateLimitType string

const (
	FixedWindow RateLimitType = "FIXED_WINDOW"
)

type FixedWindowRateLimiter struct {
	config       RateLimitConfig
	limitType    RateLimitType
	requestCount sync.Map
	windowStart  map[string]int64
	mu           sync.Mutex
}

func NewFixedWindowRateLimiter(config RateLimitConfig) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		config:      config,
		limitType:   FixedWindow,
		windowStart: make(map[string]int64),
	}
}

func (f *FixedWindowRateLimiter) AllowRequest(userId string) bool {

	allowed := false
	currentReqWindow := time.Now().Unix() / f.config.WindowInSeconds

	f.mu.Lock()
	defer f.mu.Unlock()

	countVal, _ := f.requestCount.Load(userId)
	count := 0

	if countVal != nil {
		count = countVal.(int)
	}

	lastReqWindow, exists := f.windowStart[userId]

	if !exists {
		lastReqWindow = currentReqWindow
	}

	if lastReqWindow != currentReqWindow {
		f.windowStart[userId] = currentReqWindow
		f.requestCount.Store(userId, 1)
		allowed = true
	} else if count < f.config.MaxRequests {
		f.requestCount.Store(userId, count+1)
		allowed = true
	}
	return allowed
}

func (f *FixedWindowRateLimiter) GetConfig() RateLimitConfig {
	return f.config
}

func (f *FixedWindowRateLimiter) GetType() RateLimitType {
	return f.limitType
}
