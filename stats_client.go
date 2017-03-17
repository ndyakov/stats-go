package stats

import "net/http"

// StatsClient is an interface for different methods of gathering stats
type StatsClient interface {
	// BuildTimeTracker builds timer to track metric timings
	BuildTimeTracker() TimeTracker
	// Close closes underlying client connection if any
	Close() error

	// TrackRequest tracks HTTP Request stats
	TrackRequest(r *http.Request, tt TimeTracker, success bool) StatsClient

	// TrackOperation tracks custom operation
	TrackOperation(section string, operation MetricOperation, tt TimeTracker, success bool) StatsClient
	// TrackOperation tracks custom operation with n diff
	TrackOperationN(section string, operation MetricOperation, tt TimeTracker, n int, success bool) StatsClient

	// SetHttpMetricCallback sets callback handler that allows metric operation alteration for HTTP Request
	SetHttpMetricCallback(callback HttpMetricNameAlterCallback) StatsClient

	// SetHttpRequestSection sets metric section for HTTP Request metrics
	SetHttpRequestSection(section string) StatsClient

	// ResetHttpRequestSection resets metric section for HTTP Request metrics to default value that is "request"
	ResetHttpRequestSection() StatsClient
}
