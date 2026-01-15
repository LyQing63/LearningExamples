package main

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestNewEventBus(t *testing.T) {
	bus := NewEventBus()
	if bus == nil {
		t.Fatal("NewEventBus returned nil")
	}
	if bus.handlers == nil {
		t.Fatal("NewEventBus handlers map is nil")
	}
	if bus.asyncMode != false {
		t.Fatal("NewEventBus asyncMode should be false by default")
	}
}

func TestEventBus_On_And_Emit_Sync(t *testing.T) {
	bus := NewEventBus()
	var called bool
	var receivedEvent Event

	bus.On(EventTypeHello, func(event Event) error {
		called = true
		receivedEvent = event
		return nil
	})

	evt := Event{
		ID:   1,
		Type: EventTypeHello,
		Data: "world",
	}

	err := bus.Emit(evt)
	if err != nil {
		t.Fatalf("Emit failed: %v", err)
	}

	if !called {
		t.Error("Handler was not called")
	}
	if receivedEvent.ID != 1 {
		t.Errorf("Expected event ID 1, got %d", receivedEvent.ID)
	}
	if receivedEvent.Data != "world" {
		t.Errorf("Expected event Data 'world', got %v", receivedEvent.Data)
	}
}

func TestEventBus_Emit_NoHandler(t *testing.T) {
	bus := NewEventBus()
	evt := Event{Type: EventTypeHello}
	err := bus.Emit(evt)
	if err == nil {
		t.Error("Expected error when no handler exists, got nil")
	}
}

func TestEventBus_Emit_HandlerError(t *testing.T) {
	bus := NewEventBus()
	expectedErr := errors.New("handler failed")

	bus.On(EventTypeHello, func(event Event) error {
		return expectedErr
	})

	evt := Event{Type: EventTypeHello}
	err := bus.Emit(evt)
	if err == nil {
		t.Fatal("Expected error from handler, got nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error wrapping '%v', got '%v'", expectedErr, err)
	}
}

func TestEventBus_Emit_Async(t *testing.T) {
	bus := NewEventBus()
	bus.asyncMode = true

	var wg sync.WaitGroup
	wg.Add(1)

	var receivedEvent Event

	bus.On(EventTypeHello, func(event Event) error {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond) // Simulate work
		receivedEvent = event
		return nil
	})

	evt := Event{
		ID:   2,
		Type: EventTypeHello,
	}

	err := bus.Emit(evt)
	if err != nil {
		t.Fatalf("Emit failed: %v", err)
	}

	wg.Wait()

	if receivedEvent.ID != 2 {
		t.Errorf("Expected event ID 2, got %d", receivedEvent.ID)
	}
}

func TestEventBus_Off(t *testing.T) {
	bus := NewEventBus()
	handler := func(event Event) error { return nil }

	bus.On(EventTypeHello, handler)
	if !bus.HasHandlers(EventTypeHello) {
		t.Error("Expected handlers after On")
	}

	bus.Off(EventTypeHello)
	if bus.HasHandlers(EventTypeHello) {
		t.Error("Expected no handlers after Off")
	}
}

func TestEventBus_ClearHandlers(t *testing.T) {
	bus := NewEventBus()
	bus.On(EventTypeHello, func(event Event) error { return nil })
	bus.On(EventTypeGoodbye, func(event Event) error { return nil })

	bus.ClearHandlers()

	if bus.HasHandlers(EventTypeHello) {
		t.Error("Expected no handlers for Hello after ClearHandlers")
	}
	if bus.HasHandlers(EventTypeGoodbye) {
		t.Error("Expected no handlers for Goodbye after ClearHandlers")
	}
}
