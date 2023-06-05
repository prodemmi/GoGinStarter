package event

type Listener interface {
	HandleEvent(data interface{})
}
