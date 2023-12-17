package events

type mockObserver struct {
	id          int
	lastEvent   string
	lastPayload EventPayload
}

func (m *mockObserver) OnEvent(eventName string, payload EventPayload) {
	m.lastEvent = eventName
	m.lastPayload = payload
}

func (m *mockObserver) GetID() int {
	return m.id
}
