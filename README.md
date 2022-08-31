# Go application monitoring

This codebase contains a Go server with a simple GET endpoint `/randomiser`.
This endpoint randomly returns HTTP status 200, 400 and 500, with a rough probability of 50%, 25% and 25% respectively.
The application metrics can be monitored using two observability stacks

## _Using Prometheus and Grafana_

### In order to monitor this metric using Prometheus and Grafana, follow these steps:

- Run `docker-compose -f docker-compose-prometheus.yml up -d`
- Go to `http://localhost:8000/randomiser` in the browser, or execute `curl http://localhost:8000/randomiser` in the terminal
- Access this endpoint a few times, in order to generate data points for each of the status codes being returned
- Go to `http://localhost:9090`, which is the Prometheus instance
- In the query explorer, execute metric `http_response_value` and observer the count of each of the 3 status codes returned
- Check the same metric in the Graph tab. Further filtering can be done via the labels (for example, `code`)
- Go to `http://localhost:3000` and login via the default credentials, username `admin` and password `admin`
- This Grafana instances is already integrated with the Prometheus instance. We can straightaway start creating a dashboard, using the same PromQL expression as used in Prometheus
- Using this data, alerts can be added on Grafana
- In order to stop the containers, execute `docker-compose -f docker-compose-prometheus.yml down`. This will also remove all the metric data from all the containers, since we are not persisting any data

## _Using TICK stack (Telegraf, InfluxDB, Chronograf, Kapacitor)_

### In order to monitor this metric using TICK stack, follow these steps:

- Run `docker-compose -f docker-compose-tick.yml up -d`
- Go to `http://localhost:8000/randomiser` in the browser, or execute `curl http://localhost:8000/randomiser` in the terminal
- The components within this setup are as follows:
    - The application server, in `Go`, that publishes the metric
    - The `StatsD` client library in `Go`, that is used as the medium for publishing the application metrics to `Telegraf`
    - `Telegraf`, which accumulates the metrics sent to it and stores them in `InfluxDB`
    - `InfluxDB`, acting as the time-series database for storing all the metric data
    - `Chronograf` to visualise the metrics graphically, by connecting it to the `InfluxDB` source
    - `Kapacitor`, to setup alert rules via `Chronograf` UI

    >Note: It is suggested that InfluxDB 2 (which is the latest Docker version as of this writing) can be used for alerting purposes as well, thereby not needing Kapacitor in specific ([discussion](https://community.influxdata.com/t/kapacitor-needed-for-influxdb-2-0/14862/8)). Additionally, [Kapacitor still relies on InfluxDB v1's APIs](https://github.com/influxdata/kapacitor/issues/2476). This codebase does connect Kapacitor to Chronograf anyhow

- Access this endpoint a few times, in order to generate data points for each of the status codes being returned
- Go to `http://localhost:8086`, which is the InfluxDB instance, and login via the specified credentials, username `admin` and password `admin123`
- Go to `Load data` -> `Buckets` -> `metrics` to visualise the data sent by the application. The metric to look for is `http_response_value_code_<200|400|500>`
- The `API tokens` can be used to generate a new API token to be used for interactions with InfluxDB, and can be replaced in the `.env` file in the codebase
- Go to `localhost:8888` for the Chronograf UI
- The connection with InfluxDB and Kapacitor can be checked in the `Configuration` tab
- Go to `Dashboards` tab and create a dashboard from the query explorer at the bottom of the screen. It is easier to do so if one is familiar with [`Flux`](https://docs.influxdata.com/influxdb/cloud/query-data/get-started/)
- Alerts can be set on [InfluxDB](https://docs.influxdata.com/influxdb/v2.0/monitor-alert/) or [Kapacitor via Chronograf](https://docs.influxdata.com/chronograf/v1.9/guides/create-alert-rules/)
- In order to stop the containers, execute `docker-compose -f docker-compose-tick.yml down`. This will also remove all the metric data from all the containers, since we are not persisting any data
