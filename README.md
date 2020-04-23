# GOLANG with Grafana + Prometheus Monitoring

This is a simple project to demonstrate **Golang** with **Prometheus** time series db and visualize in **Grafana**

This project using docker for service container

Do not using all code on your production !!!
# run
## [1] api
```sh
make run-go
```
## [2] docker-compose
```sh
docker-compose up
```
## [3] test
call the following url to increment the 2 counters: ``http://localhost:8080/submit``
and connect to prometheus: ``http://localhost:9090/``
or connect to grafana: ``http://localhost:3000/``


# todo
- seperate func to package instead using one package main
- do not run project as `go run main.go router.go metric-error.go metric-submit-info.go ...` it is suck!!