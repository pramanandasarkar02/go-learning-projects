package main

import (
	"context"
	"event-bus/internals"
	"fmt"
	"time"
)

func main() {
	bus := internals.NewEventBus()
	producer := internals.NewProducer(bus)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	consumer := internals.NewConsumer("c1", 2)
	bus.Subscribe(ctx, consumer)

	go func() {
		for e := range consumer.Channel() {
			fmt.Printf("Consumed offset=%d value=%s\n", e.Offset, e.Value)
			consumer.Commit(e.Offset)
			time.Sleep(200 * time.Millisecond) // slow consumer → backpressure
		}
	}()

	for i := 0; i < 5; i++ {
		producer.Send("key", []byte(fmt.Sprintf("msg-%d", i)))
	}

	time.Sleep(2 * time.Second)
}
