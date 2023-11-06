# NATS PubSub Starter Template

This example project shows a basic setup of NATS JetStream publisher / subscriber using [Watermill](https://watermill.io/). The application runs in a loop, consuming events from a NATS JetStream.

This is an example for NATS push-based consumers. For pull-based consumers, see [this branch](https://github.com/minghsu0107/NATS-PubSub/tree/pull-consumer).

There's a docker-compose file included, so you can run the example and see it in action.

## Files

- [main.go](main.go) - example source code
- [docker-compose.yml](docker-compose.yml) - local environment Docker Compose configuration
- [go.mod](go.mod) - Go modules dependencies, you can find more information at [Go wiki](https://github.com/golang/go/wiki/Modules)
- [go.sum](go.sum) - Go modules checksums

## Requirements

To run this example you will need Docker and docker-compose installed. See the [installation guide](https://docs.docker.com/compose/install/).

## Result
`subscriber1` and `subscriber2` are in the same queue group `example`, and they both subscribe to `example_topic.>`. In each round, `publisher` publishes four messages to `example_topic.a`, `example_topic.b`, `example_topic.a.test`, and `example_topic.b.test` respectively. We can see that both `subscriber1` and `subscriber2` can receive messages from all four subjects, and each message is processed only once by either `subscriber1` or `subscriber2` since they are in the same queue group.
```
> docker-compose up

2023/09/20 12:03:07 [subscriber1] received message: 0, payload: hello from a
2023/09/20 12:03:07 [subscriber1] received message: 0, payload: hello from b
2023/09/20 12:03:07 [subscriber2] received message: 0, payload: hello from a.test
2023/09/20 12:03:07 [subscriber2] received message: 0, payload: hello from b.test
2023/09/20 12:03:08 [subscriber2] received message: 1, payload: hello from a
2023/09/20 12:03:08 [subscriber2] received message: 1, payload: hello from b
2023/09/20 12:03:08 [subscriber2] received message: 1, payload: hello from a.test
2023/09/20 12:03:08 [subscriber1] received message: 1, payload: hello from b.test
2023/09/20 12:03:09 [subscriber1] received message: 2, payload: hello from a
2023/09/20 12:03:09 [subscriber1] received message: 2, payload: hello from b
2023/09/20 12:03:09 [subscriber1] received message: 2, payload: hello from a.test
2023/09/20 12:03:09 [subscriber2] received message: 2, payload: hello from b.test
2023/09/20 12:03:10 [subscriber2] received message: 3, payload: hello from a
2023/09/20 12:03:10 [subscriber1] received message: 3, payload: hello from b
2023/09/20 12:03:10 [subscriber1] received message: 3, payload: hello from a.test
2023/09/20 12:03:10 [subscriber1] received message: 3, payload: hello from b.test
...
```