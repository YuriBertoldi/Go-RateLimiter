package bd

import (
	"time"
)

type BdInterface interface {
	IncrementRequestCount(identifier string, expiry time.Duration) (int, error)
	GetRequestCount(identifier string) (int, error)
}
