package application

import (
	"time"
)

// MetricsClient allows for collecting Http Request metrics, eg. sending them to AWS CloudWatch
type MetricsClient interface {
	PutRequestMetrics(pattern string, method string, executionTime time.Duration)
}

// NullMericsClient provides a no-op implementation of MetricsClient
type NullMericsClient struct {
}

func (m *NullMericsClient) PutRequestMetrics(pattern string, method string, executionTime time.Duration) {
	// no operation
}
