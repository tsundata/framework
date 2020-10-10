package middleware

import (
	"github.com/sysatom/framework"
	"log"
	"time"
)

func Logger() framework.HandlerFunc {
	return func(c *framework.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
