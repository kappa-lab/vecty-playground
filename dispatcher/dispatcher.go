package dispatcher

import "net/url"

// ID is a unique identifier representing a registered callback function.
type ID int

var (
	idCounter ID
	callbacks = make(map[ID]func(url url.URL))
)

// Dispatch dispatches the given action to all registered callbacks.
func Dispatch(url url.URL) {
	for _, c := range callbacks {
		c(url)
	}
}

// Register registers the callback to handle dispatched actions, the returned
// ID may be used to unregister the callback later.
func Register(callback func(url url.URL)) ID {
	idCounter++
	id := idCounter
	callbacks[id] = callback
	return id
}

// Unregister unregisters the callback previously registered via a call to
// Register.
func Unregister(id ID) {
	delete(callbacks, id)
}
