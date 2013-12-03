package connect

type EventDispatcher struct {
	eventHandlers map[string][]EventHandler
}

func (this *EventDispatcher) RegisterEvent(name string, eventHandler EventHandler) {
	if this.eventHandlers == nil {
		this.eventHandlers = make(map[string][]EventHandler)
	}
	if _, ok := this.eventHandlers[name]; !ok {
		this.eventHandlers[name] = make([]EventHandler, 0)
	}
	this.eventHandlers[name] = append(this.eventHandlers[name], eventHandler)
}

func (this *EventDispatcher) DispatchEvent(name string, event Event) {
	if this.eventHandlers == nil {
		return
	}
	var eventHandlers []EventHandler
	var ok bool
	if eventHandlers, ok = this.eventHandlers[name]; !ok {
		return
	}
	for _, eventHandler := range eventHandlers {
		go eventHandler(event)
	}
}