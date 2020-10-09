package main

import (
	"net/http"
	"web"
)

func main() {
	r := web.New()

	r.GET("/", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>Home</h1>")
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

	r.GET("/assets/*filepath", func(c *web.Context) {
		c.JSON(http.StatusOK, web.H{"filepath": c.Param("filepath")})
	})

	v2 := r.Group("/v2")
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
