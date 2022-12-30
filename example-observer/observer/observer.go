package observer

// Event is an entity that is being sent by Notifiers and listened by Observers
type Event struct {
	Value interface {}
}

// Observer listens to the Events and performs OnNotify callback on every event
// that is sent to it by Notifiers
type Observer interface {
	OnNotify(*Event)
}

// Notifier collects Observers (Register / Remove) and uses Notify to ping them
// about a new event
type Notifier interface {
	Register(*Observer)
	Remove(*Observer)
	Notify(*Event)
}