package main

import (
	"log"

	"example-observer/observer"
)

type localObserver struct {
	ID int
}

type localNotifier struct {
	observers []localObserver
}

func (o *localObserver) OnNotify(e *observer.Event) {
	log.Println("Observer", o.ID, "Value", e.Value)
}

// Register adds the observer to the list of observers that Notifier tracks
func (n *localNotifier) Register(o *localObserver) {
	n.observers = append(n.observers, *o)
}

// Remove removes the Observer from the Notifier's list of tracked Observers
func (n *localNotifier) Remove(o *localObserver) {
	for i, ob := range n.observers {
		if ob == *o {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
		}
	}
}

// Notify notifies all the Observers about a new Event and calls their OnNotify
// callback
func (n *localNotifier) Notify(e *observer.Event) {
	for _, o := range n.observers {
		o.OnNotify(e)
	}
}

func main() {
	notifier := localNotifier {}

	observer1 := localObserver{ ID: 1 }
	observer2 := localObserver{ ID: 2 }

	notifier.Register(&observer1)
	notifier.Notify(&observer.Event{Value: 1})
	notifier.Register(&observer2)
	notifier.Notify(&observer.Event{Value: 2})
	notifier.Remove(&observer1)
	notifier.Notify(&observer.Event{Value: 3})
	notifier.Remove(&observer2)
	notifier.Notify(&observer.Event{Value: 4})
}