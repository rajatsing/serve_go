package main

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true // marking that the client is connected
		case message := <-manager.broadcast: // reading messages on this channel
			for conn := range manager.clients { // checking all the clients which are active
				select {
				case conn.send <- message: // Sending message
				}
			}
		}
	}
}
