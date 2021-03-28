package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//定义prometheus的metrics类型
var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device", "mount"},
	)
)

//把这个
func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

func main() {
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda", "mount": "home"}).Add(2)

	r := gin.Default()
	r.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})
	r.Run(":8002")
	//http.Handle("/metrics", promhttp.Handler())
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
