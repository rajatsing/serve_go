package main

func (manager *ClientManager) send(message []byte) {
	for conn := range manager.clients { // Looping through all the client to send messages
		conn.send <- message
	}
}
