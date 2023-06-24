package mw

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ThrottleTokenBucketMw(rLimit int, period time.Duration) fiber.Handler {
	limit := int32(rLimit)

	go func() {
		ticker := time.NewTicker(period)
		defer ticker.Stop()

		for {
			<-ticker.C
			atomic.StoreInt32(&limit, int32(rLimit))
		}
	}()

	return func(c *fiber.Ctx) error {
		fmt.Println(limit)

		if limit > 0 {
			atomic.StoreInt32(&limit, int32(limit-1))
			c.Next()
		} else {
			c.Status(http.StatusTooManyRequests)
		}

		return nil
	}
}
