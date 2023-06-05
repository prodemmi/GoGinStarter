package event

type Dispatcher interface {
	Dispatch(eventName string, eventData interface{})
}

type dispatcher struct {
	listeners map[string][]Listener
}

func (d *dispatcher) AddListener(eventName string, listener Listener) {
	if _, ok := d.listeners[eventName]; !ok {
		d.listeners[eventName] = []Listener{}
	}
	d.listeners[eventName] = append(d.listeners[eventName], listener)
}

func (d *dispatcher) Dispatch(eventName string, eventData interface{}) {
	if listeners, ok := d.listeners[eventName]; ok {
		for _, listener := range listeners {
			listener.HandleEvent(eventData)
		}
	}
}

func ProvideDispatcher() Dispatcher {
	return &dispatcher{Listeners}
}
