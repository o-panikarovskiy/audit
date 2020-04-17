package sockets

// EventHandler socket handler
type EventHandler func(*SocketClient, *SocketMessage)

var listeners map[string][]EventHandler = make(map[string][]EventHandler)

// Subscribe main func to add socket listener
func Subscribe(eventName string, handler EventHandler) {
	arr, ok := listeners[eventName]

	if !ok {
		arr = make([]EventHandler, 0)
	}

	arr = append(arr, handler)
	listeners[eventName] = arr
}

func fireEvents(conn *SocketClient, msg *SocketMessage) {
	arr, ok := listeners[msg.EventName]

	if !ok {
		return
	}

	for _, listener := range arr {
		go listener(conn, msg)
	}
}
