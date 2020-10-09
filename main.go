package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"web"
	"web/middleware"
)

func onlyForV2() web.HandlerFunc {
	return func(c *web.Context) {
		t := time.Now()
		c.Fail(http.StatusInternalServerError, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func FormatAsDate(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

func main() {
	r := web.New()
	r.Use(middleware.Logger())
	r.SetFuncMap(template.FuncMap{"FormatAsDate": FormatAsDate})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *web.Context) {
		c.String(http.StatusOK, "home")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *web.Context) {
			c.String(http.StatusOK, "v1 Hi %s, %s", c.Query("name"), c.Path)
		})

		v1.GET("/hello/:name", func(c *web.Context) {
			c.String(http.StatusOK, "Hi %s, %s", c.Param("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v1.GET("/hello", func(c *web.Context) {
			c.String(http.StatusOK, "v2 Hi %s, %s", c.Query("name"), c.Path)
		})
		v2.POST("/login", func(c *web.Context) {
			c.JSON(http.StatusOK, web.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":5000")
}
