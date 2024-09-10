package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var DurationSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "Duration",
	Name:      "http_server_request_duration_seconds",
	Help:      "Histogram of response time for handler in seconds",
	Buckets:   []float64{.001, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{"route"})

var ResponseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"route", "status"})

var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"route"},
)
