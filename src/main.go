package main

import (
  "log"
  "net/http"
  "strconv"

  "github.com/prometheus/client_golang/prometheus/promhttp"

  "go-monitoring/service"
)

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
}


func main() {
    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/randomiser", randomResponseGenerator)
    log.Println("Listening for requests at http://localhost:8000/randomiser")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
