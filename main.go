package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var (
	listenAddr string
)

func main() {

	flag.StringVar(&listenAddr, "addr", ":8087", "web的addr")

	flag.Parse()

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/hello/:user", helloUser)
	router.GET("/code/:token", multiStatusCode)

	router.Run(listenAddr)
}

// 直接返回传入的user，体现在日志中
// [GIN] 2022/05/13 - 17:02:40 | 200 |            0s |             ::1 | GET      "/hello/xiaoyi"
func helloUser(c *gin.Context) {

	user := c.Param("user")

	c.String(200, user)
}

// 根据传入的token返回不同的状态码，体现在日志中
// [GIN] 2022/05/13 - 17:01:17 | 200 |            0s |             ::1 | GET      "/code/a1"
// [GIN] 2022/05/13 - 17:01:26 | 300 |            0s |             ::1 | GET      "/code/b1"
func multiStatusCode(c *gin.Context) {

	token := c.Param("token")
	var rc int
	switch token {
	case "a1":
		rc = 200
	case "b1":
		rc = 300
	case "c1":
		rc = 400
	case "d1":
		rc = 500

	default:
		rc = 201

	}
	c.String(rc, token)
}
