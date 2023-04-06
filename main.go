package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func main() {

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hola, Mundo!!!")
	})

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Run()

}
