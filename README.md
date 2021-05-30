# NATS PubSub Starter Template

This example project shows a basic setup of NATS Streaming publisher / subscriber using [Watermill](https://watermill.io/). The application runs in a loop, consuming events from a NATS Streaming cluster.

There's a docker-compose file included, so you can run the example and see it in action.

## Files

- [main.go](main.go) - example source code
- [docker-compose.yml](docker-compose.yml) - local environment Docker Compose configuration, contains Golang, Kafka and Zookeeper
- [go.mod](go.mod) - Go modules dependencies, you can find more information at [Go wiki](https://github.com/golang/go/wiki/Modules)
- [go.sum](go.sum) - Go modules checksums

## Requirements

To run this example you will need Docker and docker-compose installed. See the [installation guide](https://docs.docker.com/compose/install/).

## Usage

```bash
> docker-compose up

2021/05/30 04:55:40 received message: 2154756c-e621-41d2-bc0a-c3a8ec11b19d, payload: Hello, world!
2021/05/30 04:55:41 received message: 3d9ddc61-e4fc-47a0-a245-7c44c81614e3, payload: Hello, world!
2021/05/30 04:55:42 received message: 5e1b9f37-d964-4520-b1fc-a74e55404ac3, payload: Hello, world!
...
```