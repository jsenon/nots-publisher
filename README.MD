[![CircleCI](https://circleci.com/gh/jsenon/nats-publisher.svg?style=svg)](https://circleci.com/gh/jsenon/nats-publisher)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsenon/nats-publisher)](https://goreportcard.com/report/github.com/jsenon/nats-publisher)

# Nats Publisher
 
This is a Simple nats publisher

### Func

Use container in order to debug nats 

## Build

Build locally
```
make 
```

## Run Container

Run Container
```
docker run -d -e MY_NATSBOOTSTRAP="nats://localhost:4222"  jsenon/nats-publisher:latest
```

## Run on K8S Cluster

Deploy POD
```
kubectl apply -f k8s/deployment.yml
```

## Env


## Todo

