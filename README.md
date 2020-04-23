# GOLANG with Grafana + Prometheus Monitoring

This is a simple project to demonstrate **Golang** with **Prometheus** time series db and visualize in **Grafana**

This project using docker for service container

Do not using all code on your production !!!

# configure
Replace `REPLACE_WITH_LOCAL_IP` with for example `192.168.1.2` in `prometheus/prometheus.yml`

# run
## [1] api
```sh
make run-go
```
## [2] docker-compose
```sh
docker-compose up
```

# todo
- seperate func to package instead using one package main
- do not run project as `go run main.go router.go metric-error.go metric-submit-info.go ...` it is suck!!