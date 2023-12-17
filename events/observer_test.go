package events

import "testing"

func TestEventPublisher_AddRemoveListeners(t *testing.T) {
	publisher := NewEventPublisher()

	observer1 := &mockObserver{
		id: 1,
	}
	observer2 := &mockObserver{
		id: 2,
	}

	publisher.AddListener(observer1, "event1")
	publisher.AddListener(observer1, "event1") // registering to the same event twice
	publisher.AddListener(observer2, "event1")
	publisher.AddListener(observer2, "event2")

	if len(publisher.eventObservers["event1"]) != 2 {
		t.Error("expected only 2 listeners on event1")
	}
	if len(publisher.eventObservers["event2"]) != 1 {
		t.Error("expected only 1 listener on event2")
	}

	// removing non existent listener
	publisher.RemoveListener(observer1, "event2")
	if len(publisher.eventObservers["event2"]) != 1 {
		t.Error("expected only 1 listener on event2")
	}

	publisher.RemoveListener(observer1, "event1")
	publisher.RemoveListener(observer2, "event1")
	if len(publisher.eventObservers["event1"]) != 0 {
		t.Error("expected no listeners on event1")
	}
}

func TestEventPublisher_FireEvent(t *testing.T) {
	publisher := NewEventPublisher()

	observer1 := &mockObserver{
		id: 1,
	}
	observer2 := &mockObserver{
		id: 2,
	}

	publisher.AddListener(observer1, "event1")
	publisher.AddListener(observer2, "event1")
	publisher.AddListener(observer2, "event2")

	payload1 := EventPayload{"pay1": "val1"}
	payload2 := EventPayload{"pay2": "val2"}

	publisher.FireEvent("event1", payload1)
	if observer1.lastEvent != "event1" || observer1.lastPayload["pay1"] != "val1" {
		t.Error("observer missed event[1]")
	}
	if observer2.lastEvent != "event1" || observer2.lastPayload["pay1"] != "val1" {
		t.Error("observer missed event[2]")
	}

	publisher.FireEvent("event2", payload2)
	if observer1.lastEvent != "event1" || observer1.lastPayload["pay1"] != "val1" {
		t.Error("observer missed event[3]")
	}
	if observer2.lastEvent != "event2" || observer2.lastPayload["pay2"] != "val2" {
		t.Error("observer missed event[4]")
	}
}
