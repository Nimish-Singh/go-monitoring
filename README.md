# Go application monitoring
## _Using Prometheus and Grafana_

This codebase contains a Go server with a simple GET endpoint `/randomiser`.
This endpoint randomly returns HTTP status 200, 400 and 500, with a rough probability of 50%, 25% and 25% respectively.

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
