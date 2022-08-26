package metrics

import (
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    Http_response = promauto.NewCounterVec(prometheus.CounterOpts{
      Name: "http_response_value",
      Help: "The HTTP response for a request",
    },
    []string{"code", "url"},
  )
)
