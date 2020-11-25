package metric

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

type Interceptor struct {
	m *Metric
}

func NewInterceptor(m *Metric) *Interceptor {
	return &Interceptor{
		m: m,
	}
}

func (i *Interceptor) Collect(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	t := time.Now()

	resp, err := handler(ctx, req)
	result := "OK"
	if err != nil {
		result = err.Error()
	}
	i.m.TotalHits.Inc()
	i.m.Hits.WithLabelValues(result, info.FullMethod, info.FullMethod).Inc()
	i.m.Durations.WithLabelValues(result, info.FullMethod, info.FullMethod).Observe(time.Since(t).Seconds())
	return resp, err
}
