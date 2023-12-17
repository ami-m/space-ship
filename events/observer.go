package events

type EventPayload map[string]any

// Observer is a receiver of events
type Observer interface {
	OnEvent(eventName string, payload EventPayload)
	GetID() int
}

// Subject is a publisher of events
type Subject interface {
	AddListener(observer Observer, eventName string)
	RemoveListener(observer Observer, eventName string)
	FireEvent(eventName string, payload EventPayload)
}

type EventPublisher struct {
	eventObservers map[string]map[int]Observer // eventName -> observer (by observerId)
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{make(map[string]map[int]Observer)}
}

func (p *EventPublisher) AddListener(observer Observer, eventName string) {
	_, ok := p.eventObservers[eventName]
	if !ok {
		p.eventObservers[eventName] = make(map[int]Observer)
	}
	p.eventObservers[eventName][observer.GetID()] = observer
}

func (p *EventPublisher) RemoveListener(observer Observer, eventName string) {
	observers, ok := p.eventObservers[eventName]
	if !ok {
		return
	}
	delete(observers, observer.GetID())
}

func (p *EventPublisher) FireEvent(eventName string, payload EventPayload) {
	observers, ok := p.eventObservers[eventName]
	if !ok {
		return
	}
	for _, observer := range observers {
		observer.OnEvent(eventName, payload)
	}
}
