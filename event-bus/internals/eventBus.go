package internals

import (
	"context"
	"sync"
	"time"
)

type EventBus struct {
	log       *EventLog
	consumers map[string]*Consumer
	mu        sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		log:       NewEventLog(),
		consumers: make(map[string]*Consumer),
	}
}


func (b *EventBus) Publish(key string, value []byte) {
	event := Event{Key: key, Value: value}
	b.log.Append(event)
}


func (b *EventBus) Subscribe(ctx context.Context, c *Consumer) {
	b.mu.Lock()
	b.consumers[c.ID] = c
	b.mu.Unlock()

	go func() {
		defer close(c.ch)

		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				events := b.log.ReadFrom(c.offset)
				for _, e := range events {
					select {
					case c.ch <- e: // backpressure here
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()
}
