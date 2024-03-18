package limiter

type RateLimitInterface interface {
	IsLimitExceeded(identifier string) (bool, error)
}
