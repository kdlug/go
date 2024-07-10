// We are implementing notifier interface using a pointer semantics
// but we are passing a value of type user -> func sentNotification(n notifier)
package main

import "fmt"

type user struct {
	name  string
	email string
}

type notifier interface {
	notify()
}

// notify implements the notifier interface
func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>",
		u.name,
		u.email)
}

// sendNotification accepts values that implement
// the notifier interface and sends notifications
func sentNotification(n notifier) {
	n.notify()
}

func main() {
	// create value of type user
	u := user{"Pablo", "pablo@picasso.com"}

	// values of type user do not implement the interface
	// because pointer reveivers don't belong to the metod set of a value
	// sentNotification(u) // value semantic

	// we have to pass the pointer
	sentNotification(&u)
}
