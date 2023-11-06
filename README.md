# NATS PubSub Starter Template (Pull-based Consumer)

This example project shows a basic setup of NATS JetStream publisher / subscriber using [Watermill](https://watermill.io/). The application runs in a loop, consuming events from a NATS JetStream.

There's a docker-compose file included, so you can run the example and see it in action.

## Files

- [main.go](main.go) - example source code
- [docker-compose.yml](docker-compose.yml) - local environment Docker Compose configuration
- [go.mod](go.mod) - Go modules dependencies, you can find more information at [Go wiki](https://github.com/golang/go/wiki/Modules)
- [go.sum](go.sum) - Go modules checksums

## Requirements

To run this example you will need Docker and docker-compose installed. See the [installation guide](https://docs.docker.com/compose/install/).

## Result
`subscriber1` and `subscriber2` represent two subscriptions bound to the same consumer `myconsumer`, and they both subscribe to `example_topic.test`. `publisher` publishes a messages to `example_topic.test` every 50 millisecond. We can see that a message is processed by either `subscriber1` or `subscriber2` once since the stream `example_topic` uses the work-queue retention policy.