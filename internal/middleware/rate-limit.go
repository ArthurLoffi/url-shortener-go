package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type IpRateLimiter struct {
	mu sync.Mutex
	limiters map[string]*rate.Limiter
}

func NewIPRateLimiter() *IpRateLimiter {
	return &IpRateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

func (i *IpRateLimiter) getLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.limiters[ip]

	if !exists {
		// 5 requisições por segundo com burst de 10
		limiter = rate.NewLimiter(5, 10)
		i.limiters[ip] = limiter
	}

	return limiter
}

func RateLimitMiddleware(limiter *IpRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.ClientIP()

		if !limiter.getLimiter(ip).Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}