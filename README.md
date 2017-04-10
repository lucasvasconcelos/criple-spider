# criple-spider
A _criple_ web crawler. Use this to demonstrate distributed applications on Kubernetes on presentations.

# Build binary
`$ make`

# Build image
`$ make container`

# Push to registry
`$ make push`

# Running and seein the metrics

- `$ docker-compose up -d`
- Open Grafana on http://localhost:3000 (admin:secret)
- Add a *InfluxDB* datasource named `criple-spider` and make it as default, using this address: http://influxdb-criple-spider:8086
- Use a database named `criple-spider`
- Import the dashboard found in `grafana/criple-spider-dashboard.json` and select the datasource previously added 
- The dashboard will show the visited pages in the last 1 minute.

