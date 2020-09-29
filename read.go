package main

import "encoding/json"

// Read GoRoutine, its function is to "read" the data from the client
func (c *Client) read() {
	for {
		_, message, err := c.socket.ReadMessage() // Reading message here
		if err != nil {                           // If there's some issue reading the message
			c.socket.Close() // Close that socket for that client
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.name, Content: string(message)}) // marshalling the basic data type to JSON data
		manager.broadcast <- jsonMessage                                                   //in order to broadcast that json data that we marshalled, we will send it to boradcast channel
	}
}
