package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(DurationSeconds())
	prometheus.MustRegister(ResponseStatus())

}
func DurationSeconds() *prometheus.HistogramVec {

	buckets := []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}

	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "namespace",
		Name:      "http_server_request_duration_seconds_sum",
		Help:      "Histogram of response time for handler in seconds",
		Buckets:   buckets,
	}, []string{"route", "method", "status_code"})

}

func ResponseStatus() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_status",
			Help: "Status of HTTP response",
		},
		[]string{"route", "method", "status_code"},
	)

}
