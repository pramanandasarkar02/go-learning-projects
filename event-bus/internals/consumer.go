package internals

import (
	
)

type Consumer struct {
	ID     string
	offset uint64
	ch     chan Event
}

func NewConsumer(id string, buffer int) *Consumer {
	return &Consumer{
		ID: id,
		ch: make(chan Event, buffer),
	}
}

func (c *Consumer) Channel() <-chan Event {
	return c.ch
}

func (c *Consumer) Commit(offset uint64) {
	c.offset = offset + 1
}
