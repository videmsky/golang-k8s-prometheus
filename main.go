package main

import (
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hola, Mundo!!!")
		recordMetrics()
	})

	r.GET("/metrics", prometheusHandler())

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Run(":8889")

}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gkp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
