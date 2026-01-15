package main

import (
	"errors"
	"testing"
)

// MockPlugin ensures we can test the execution order and behavior
type MockPlugin struct {
	events   []EventType
	id       string
	executed bool
	lastData any
	logic    func(data any, next func() error) error
}

func (p *MockPlugin) OnEvent(eventType EventType, data any, next func() error) error {
	p.executed = true
	p.lastData = data
	if p.logic != nil {
		return p.logic(data, next)
	}
	return next()
}

func (p *MockPlugin) ActivationEvents() []EventType {
	return p.events
}

func TestPiplineManager(t *testing.T) {
	// Helper to create a manager with initialized maps
	newManager := func() *PiplineManager {
		return &PiplineManager{
			listeners: make(map[EventType][]Plugin),
			handlers:  make(map[EventType]func(EventType, any) error),
		}
	}

	t.Run("Test Registration and Execution Order", func(t *testing.T) {
		manager := newManager()
		executionOrder := []string{}

		p1 := &MockPlugin{
			id:     "p1",
			events: []EventType{EventTypeHello},
			logic: func(data any, next func() error) error {
				executionOrder = append(executionOrder, "p1-start")
				err := next()
				executionOrder = append(executionOrder, "p1-end")
				return err
			},
		}

		p2 := &MockPlugin{
			id:     "p2",
			events: []EventType{EventTypeHello},
			logic: func(data any, next func() error) error {
				executionOrder = append(executionOrder, "p2-start")
				err := next()
				executionOrder = append(executionOrder, "p2-end")
				return err
			},
		}

		manager.Register(p1)
		manager.Register(p2)

		err := manager.Trigger(EventTypeHello, "world")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Expected order: p1-start -> p2-start -> p2-end -> p1-end
		expected := []string{"p1-start", "p2-start", "p2-end", "p1-end"}
		if len(executionOrder) != len(expected) {
			t.Errorf("expected execution length %d, got %d", len(expected), len(executionOrder))
		}
		for i, val := range expected {
			if i >= len(executionOrder) {
				break
			}
			if executionOrder[i] != val {
				t.Errorf("expected %s at index %d, got %s", val, i, executionOrder[i])
			}
		}
	})

	t.Run("Test Event Filtering", func(t *testing.T) {
		manager := newManager()
		p1 := &MockPlugin{events: []EventType{EventTypeHello}}
		p2 := &MockPlugin{events: []EventType{EventTypeGoodbye}}

		manager.Register(p1)
		manager.Register(p2)

		// Trigger Hello
		manager.Trigger(EventTypeHello, nil)
		if !p1.executed {
			t.Error("p1 should be executed on Hello")
		}
		if p2.executed {
			t.Error("p2 should NOT be executed on Hello")
		}

		// Reset
		p1.executed = false
		p2.executed = false

		// Trigger Goodbye
		manager.Trigger(EventTypeGoodbye, nil)
		if p1.executed {
			t.Error("p1 should NOT be executed on Goodbye")
		}
		if !p2.executed {
			t.Error("p2 should be executed on Goodbye")
		}
	})

	t.Run("Test Chain Interruption", func(t *testing.T) {
		manager := newManager()
		p1 := &MockPlugin{
			events: []EventType{EventTypeHello},
			logic: func(data any, next func() error) error {
				// Don't call next()
				return errors.New("stop here")
			},
		}
		p2 := &MockPlugin{
			events: []EventType{EventTypeHello},
			logic: func(data any, next func() error) error {
				return next()
			},
		}

		manager.Register(p1)
		manager.Register(p2)

		err := manager.Trigger(EventTypeHello, nil)
		if err == nil {
			t.Error("expected error from p1")
		}
		if err.Error() != "stop here" {
			t.Errorf("expected 'stop here' error, got %v", err)
		}
		if p2.executed {
			t.Error("p2 should NOT be executed if p1 does not call next()")
		}
	})

	t.Run("Test Data Passing", func(t *testing.T) {
		manager := newManager()
		expectedData := "some data"

		p1 := &MockPlugin{
			events: []EventType{EventTypeHello},
		}

		manager.Register(p1)
		manager.Trigger(EventTypeHello, expectedData)

		if p1.lastData != expectedData {
			t.Errorf("expected data %v, got %v", expectedData, p1.lastData)
		}
	})
}
