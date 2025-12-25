package internals

import "sync"

type EventLog struct {
	mu     sync.RWMutex
	events []Event
}

func NewEventLog() *EventLog {
	return &EventLog{}
}

func (l *EventLog) Append(e Event) uint64 {
	l.mu.Lock()
	defer l.mu.Unlock()

	e.Offset = uint64(len(l.events))
	l.events = append(l.events, e)
	return e.Offset
}

func (l *EventLog) ReadFrom(offset uint64) []Event {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if offset >= uint64(len(l.events)) {
		return nil
	}
	return append([]Event(nil), l.events[offset:]...)
}

func (l *EventLog) Size() uint64 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return uint64(len(l.events))
}
