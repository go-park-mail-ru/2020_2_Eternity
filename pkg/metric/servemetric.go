package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
)

func RouterForMetrics(address string) {
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := r.RunListener(lis); err != nil {
		log.Fatal(err)
		return
	}
}
