package mw

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FixedWindowMw(rLimit int, restoreEvery time.Duration) fiber.Handler {
	counter := int32(0)

	go func() {
		ticker := time.NewTicker(restoreEvery)
		defer ticker.Stop()

		for {
			<-ticker.C
			atomic.StoreInt32(&counter, int32(0))
		}
	}()

	return func(c *fiber.Ctx) error {
		if atomic.LoadInt32(&counter) <= int32(rLimit) {
			atomic.AddInt32(&counter, 1)
			fmt.Println(counter)
			c.Next()
		} else {
			c.Status(http.StatusTooManyRequests)
		}

		return nil
	}
}
