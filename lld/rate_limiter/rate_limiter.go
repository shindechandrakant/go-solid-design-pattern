package rate_limiter

type RateLimitConfig struct {
	WindowInSeconds int64
	MaxRequests     int
}

type RateLimiter interface {
	AllowRequest(UserId string) bool
	GetConfig() RateLimitConfig
	GetType() RateLimitType
}
