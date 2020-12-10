package metric

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CollectMetrics(m *Metric) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		m.TotalHits.Inc()
		if c.Writer.Status() != http.StatusOK {
			m.Errors.WithLabelValues(strconv.Itoa(c.Writer.Status()), c.Request.Method, c.FullPath()).Inc()
		} else {
			m.Hits.WithLabelValues(strconv.Itoa(c.Writer.Status()), c.Request.Method, c.FullPath()).Inc()
		}
		m.Durations.WithLabelValues(strconv.Itoa(c.Writer.Status()), c.Request.Method, c.FullPath()).Observe(time.Since(t).Seconds())
	}
}
