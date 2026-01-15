package main

import (
	"fmt"
	"sync"
)

// Event 作为事件内容主题的结构体
type Event struct {
	ID       uint8
	Data     any
	MetaData any
	Type     EventType
}

// EventType 作为事件类型的枚举
type EventType string

const (
	// EventTypeHello 作为事件类型的枚举
	EventTypeHello EventType = "hello"
	// EventTypeGoodbye 作为事件类型的枚举
	EventTypeGoodbye EventType = "goodbye"
)

type EventHandler func(event Event) error

// EventBus 作为事件控制的核心
type EventBus struct {
	// 事件的读写锁，用于发布订阅的读多写少的业务
	mu sync.RWMutex
	// 事件处理器映射
	handlers map[EventType][]EventHandler
	// 是否开启同步处理模式，默认关闭，开启后，多个Handler并发处理事件时，会同步处理
	asyncMode bool
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers:  make(map[EventType][]EventHandler),
		asyncMode: false,
	}
}

// 注册动作
func (e *EventBus) On(eventType EventType, handler EventHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.handlers[eventType] = append(e.handlers[eventType], handler)
}

// 注销事件
func (e *EventBus) Off(eventType EventType) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.handlers, eventType)
}

// 发布事件
func (e *EventBus) Emit(event Event) error {
	eventType := event.Type

	e.mu.RLock()
	handlers := e.handlers[eventType]
	e.mu.RUnlock()

	if len(handlers) == 0 {
		return fmt.Errorf("this eventType %s has no any handler", event.Type)
	}

	// 异步模式处理
	if e.asyncMode {
		for _, handler := range handlers {
			h := handler
			go func() {
				_ = h(event)
			}()
		}
		return nil
	}
	// 同步模式处理
	for _, handler := range handlers {
		// 采用快速失败的方式处理
		err := handler(event)
		if err != nil {
			return fmt.Errorf("handler eventType %s has error %w", event.Type, err)
		}
	}
	return nil
}

// 是否有handlers
func (e *EventBus) HasHandlers(eventType EventType) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	_, ok := e.handlers[eventType]
	return ok
}

// 清空handlers
func (e *EventBus) ClearHandlers() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.handlers = make(map[EventType][]EventHandler)
}
