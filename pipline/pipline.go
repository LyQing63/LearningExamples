package main

type EventType string

const (
	EventTypeHello   EventType = "hello"
	EventTypeGoodbye EventType = "goodbye"
)

// 作为pipline每个节点的handler
type Plugin interface {
	OnEvent(eventType EventType, data any, next func() error) error
	// 可以注册多个事件类型
	ActivationEvents() []EventType
}

// pipline的管理器
type PiplineManager struct {
	listeners map[EventType][]Plugin
	handlers  map[EventType]func(EventType, any) error
}

func (m *PiplineManager) Register(plugin Plugin) {
	for _, eventType := range plugin.ActivationEvents() {
		m.listeners[eventType] = append(m.listeners[eventType], plugin)
		// 重新构建处理链
		m.handlers[eventType] = m.buildHandler(m.listeners[eventType])
	}
}

// 洋葱模型
func (m *PiplineManager) buildHandler(plugins []Plugin) func(EventType, any) error {
	// 末尾结束处理器
	next := func(EventType, any) error {
		return nil
	}
	// 遍历所有处理器从后往前，栈式的
	for i := len(plugins) - 1; i >= 0; i-- {
		current := plugins[i]
		preNext := next
		next = func(eventType EventType, data any) error {
			return current.OnEvent(eventType, data, func() error {
				return preNext(eventType, data)
			})
		}
	}

	return next
}

// 触发事件
func (m *PiplineManager) Trigger(eventType EventType, data any) error {
	if handler, ok := m.handlers[eventType]; ok {
		return handler(eventType, data)
	}
	return nil
}
