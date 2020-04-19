package sockets

// EventHandler socket handler
type EventHandler func(ISocketClient, *SocketMessage)

var listeners map[string][]EventHandler = make(map[string][]EventHandler)

// Subscribe func to add socket listener
func Subscribe(eventName string, handler EventHandler) {
	arr, ok := listeners[eventName]

	if !ok {
		arr = make([]EventHandler, 0)
	}

	arr = append(arr, handler)
	listeners[eventName] = arr
}

// SubscribeEvents func to add socket listeners
func SubscribeEvents(events SocketHandlers) {
	for _, val := range *events {
		Subscribe(val.Event, val.Handler)
	}
}

func fireEvents(client ISocketClient, msg *SocketMessage) {
	arr, ok := listeners[msg.EventName]

	if !ok {
		return
	}

	for _, listener := range arr {
		go listener(client, msg)
	}
}
