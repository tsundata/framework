package middleware

import (
	"log"
	"time"
	"web"
)

func Logger() web.HandlerFunc {
	return func(c *web.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
