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
docker run -d jsenon/nats-publisher:latest
```

## Run on K8S Cluster

Deploy POD
```
kubectl apply -f k8s/deployment.yml
```

## Usage

If you don't want to expose your producer externaly you can use port forward
```
kubectl port-forward nats-publisher-xxx 9010:9010 &
```

Use PostMan or Curl to send information to your producer
```
curl -X POST \
  http://127.0.0.1:9010/publish \
  -d '{
	"nats":"nats://USER:PWD@nats:4222",
	"subject":"kikoo",
	"message":"Hello World"
}'
```

Where:

* nats: your nats server
* suject: your topic
* message: your message to sent


