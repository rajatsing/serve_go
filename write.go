package main

import "github.com/gorilla/websocket"

// No we have gotten the data by reading , now lets write it
func (c *Client) write() {
	for {
		message, ok := <-c.send
		if !ok { // error handling incase, unable to read the message
			return
		}
		c.socket.WriteMessage(websocket.TextMessage, message) // write that message to the client
	}
}
