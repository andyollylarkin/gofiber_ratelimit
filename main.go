package main

import (
	"time"

	mw "github.com/andyollylarkin/gofiber_ratelimit/middlewares"
	"github.com/gofiber/fiber/v2"
)

func NotThrottledRoute(r *fiber.Ctx) error {
	r.Accepts("application/json")

	return r.SendString("Hello")
}

func ThrottledRoute(r *fiber.Ctx) error {
	r.Accepts("application/json")

	return r.SendString("Hello throttled")
}

func main() {
	app := fiber.New()
	mw := mw.ThrottleTokenBucketMw(10, time.Second*1)
	nt := app.Group("/not_throttled")
	t := app.Group("/throttled", mw)

	nt.Get("/get", NotThrottledRoute)
	t.Get("/get", ThrottledRoute)
	app.Listen(":8080")
}
