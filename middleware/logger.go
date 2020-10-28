package middleware

import (
	"log"
	"time"

	"github.com/tsundata/framework"
)

func Logger() framework.HandlerFunc {
	return func(c *framework.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
