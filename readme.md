# mahasiswa-microservice
product-microservice

# How to run

- Set up .env
```
cp .env.example .env
```

- Consul
```
consul agent -dev
```

- Promtheus Container (docker)

  Start a Prometheus container
To start collecting metrics, you need to have Prometheus running. The easiest way to do it is by using Docker and running the following command in the root folder of this repository (so it uses the prometheus.yaml config file). Replace '192.168.1.60' with your local private IP:

```
docker run --name prometheus --add-host="localhost:192.168.1.60" -d -p 9090:9090 -v ${PWD}/prometheus.yaml:/etc/prometheus/prometheus.yml prom/prometheus
```

Open your browser and go to http://localhost:9090/

- Zipkin Container (docker)
  
  To collect the traces of the applicacion, you need to have Zipkin running. Again, the easiest way to do it is by using Docker and running the following command:

```
use go.opencensus.io v0.20.1
```

```
docker run -d -p 9411:9411 openzipkin/zipkin
```

Open your browser and go to http://localhost:9411/

- server
```
go run main.go
```