package metric

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func CollectMetrics(m *Metric) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		m.TotalHits.Inc()
		m.Hits.WithLabelValues(strconv.Itoa(c.Writer.Status()), c.Request.Method, c.FullPath()).Inc()
		m.Durations.WithLabelValues(strconv.Itoa(c.Writer.Status()), c.Request.Method, c.FullPath()).Observe(time.Since(t).Seconds())
	}
}
