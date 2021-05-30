package main

import (
	"context"
	"log"
	"time"

	stan "github.com/nats-io/stan.go"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	subscriber, err := nats.NewStreamingSubscriber(
		nats.StreamingSubscriberConfig{
			ClusterID: "test-cluster",
			ClientID:  "example-subscriber",

			// When QueueGroup is empty, subscribe without QueueGroup will be used
			// in this case, SubscribersCount should be set to 1 to avoid duplication
			QueueGroup:       "example",
			DurableName:      "my-durable",
			SubscribersCount: 4, // how many goroutines should consume messages
			CloseTimeout:     time.Minute,
			// How long subscriber should wait for Ack/Nack. When no Ack/Nack was received, message will be redelivered.
			// It is mapped to stan.AckWait option
			AckWaitTimeout: time.Second * 30,
			StanOptions: []stan.Option{
				stan.NatsURL("nats://nats-streaming:4222"),
			},
			Unmarshaler: nats.GobMarshaler{},
			StanSubscriptionOptions: []stan.SubscriptionOption{
				// Read from the beginning of the channel
				// stan.DeliverAllAvailable()

				// read from the last message received
				// stan.StartWithLastReceived()

				// Read from a specific time the message arrived in the channel
				// stan.StartAtTime(startTime)

				// set the maximum number of messages the cluster will send without an ACK
				// just like RabbitMQ prefetch
				// if MaxInflight is too large, the subscriber will have not enough time processing messages
				// before NATS timeout. In this case, NATS will send the same batch of messages
				// to another subscriber in the same queue group. Thus, messages may be processed twice
				stan.MaxInflight(2048), // default: 1024
			},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}
	messages, err := subscriber.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}
	go process(messages)

	publisher, err := nats.NewStreamingPublisher(
		nats.StreamingPublisherConfig{
			ClusterID: "test-cluster",
			ClientID:  "example-publisher",
			StanOptions: []stan.Option{
				stan.NatsURL("nats://nats-streaming:4222"),
				// MaxPubAcksInflight is an Option to set the maximum number of published
				// messages without outstanding ACKs from the server
				// default: 16384
				stan.MaxPubAcksInflight(18000),
			},
			Marshaler: nats.GobMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	publishMessages(publisher)
}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))

		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
