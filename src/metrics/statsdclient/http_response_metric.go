package statsdclient

import (
  "log"

  "github.com/cactus/go-statsd-client/v5/statsd"
)

var Client = getClient()

func getClient() statsd.Statter {
  config := &statsd.ClientConfig{
        Address: "host.docker.internal:8125",
        Prefix: "",
        TagFormat: statsd.InfixSemicolon,
    }

  client, err := statsd.NewClientWithConfig(config)

  if err != nil {
     log.Printf("Error in initialising statsd client- %v", err)
     return nil
  }

  return client
}
