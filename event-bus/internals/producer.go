package internals

type Producer struct {
	bus *EventBus
}

func NewProducer(bus *EventBus) *Producer {
	return &Producer{bus: bus}
}

func (p *Producer) Send(key string, value []byte) {
	p.bus.Publish(key, value)
}
