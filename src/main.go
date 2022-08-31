package main

import (
  "log"
  "net/http"
  "os"
  "strconv"
  "strings"

  "github.com/cactus/go-statsd-client/v5/statsd"
  "github.com/prometheus/client_golang/prometheus/promhttp"

  "go-monitoring/metrics/prometheusclient"
  "go-monitoring/metrics/statsdclient"
  "go-monitoring/service"
)

const (
  prometheus = "prometheus"
  tick = "tick"
)

func collectMetric(responseCode string, url string) {
  var observabilityStack = strings.ToLower(os.Getenv("OBSERVABILITY_STACK"))
  if observabilityStack == tick {
    log.Print("tick stack")
    err := statsdclient.Client.Inc(fmt.Sprintf("http_response_value_code_%s", responseCode), 1, 1.0, statsd.Tag{"url", url})
    if err != nil {
      log.Printf("Could not send metrics. Error: %v", err)
    }
    return
  }
  if observabilityStack == prometheus {
    log.Print("prom stack")
    prometheusclient.Http_response.WithLabelValues(responseCode, url).Inc()
  }
}

func getResponseCode() int {
  randomNumber := service.Generate()
  var status int

  if randomNumber <= 50 {
    status = http.StatusOK
  } else if randomNumber <= 75 {
    status = http.StatusBadRequest
  } else {
    status = http.StatusInternalServerError
  }
  return status
}

func randomResponseGenerator (w http.ResponseWriter, req *http.Request) {
    responseCode := getResponseCode()
    log.Printf("Returning status code- %d", responseCode)
    w.WriteHeader(responseCode)
    w.Write([]byte("Returned status code:"  + strconv.Itoa(responseCode)))
    collectMetric(strconv.Itoa(responseCode), "/randomiser")
}


func main() {
    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/randomiser", randomResponseGenerator)
    log.Println("Listening for requests at http://localhost:8000/randomiser")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
