package mw

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ThrottleTokenBucketMw(rLimit int, refillEvery time.Duration) fiber.Handler {
	tokens := int32(rLimit)

	go func() {
		ticker := time.NewTicker(refillEvery)
		defer ticker.Stop()

		for {
			<-ticker.C

			if tokens < int32(rLimit) {
				atomic.StoreInt32(&tokens, int32(tokens+1))
			}
			fmt.Println(tokens)
		}
	}()

	return func(c *fiber.Ctx) error {
		fmt.Println(tokens)

		if tokens > 0 {
			atomic.StoreInt32(&tokens, int32(tokens-1))
			c.Next()
		} else {
			c.Status(http.StatusTooManyRequests)
		}

		return nil
	}
}
