package main

import (
	"github.com/ATao13/gee_web/day2-context/gee"
	"net/http"
)

func main() {
	r := gee.New()
	//注册路由，并映射处理方法
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostFrom("username"),
			"password": c.PostFrom("password")})
	})
	// r 实现了Handler接口
	r.Run(":6666")
}
