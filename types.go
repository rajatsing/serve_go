package main

import "github.com/gorilla/websocket"

// This will keep track of all connected clients
var manager = ClientManager{
	broadcast: make(chan []byte),
	register:  make(chan *Client),
	clients:   make(map[*Client]bool),
}

// keeps track of all clients
type ClientManager struct {
	clients   map[*Client]bool // client struct
	broadcast chan []byte      // message that needs to be broadcasted
	register  chan *Client     // registred clients
}

// This is the actual "client" struct
type Client struct {
	name   string          // unique name
	socket *websocket.Conn // socket Connection details
	send   chan []byte     // message to be sent
}

//  JSON that we will use to pass the data
type Message struct {
	Sender  string `json:"sender,omitempty"`  // sender name
	Content string `json:"content,omitempty"` // content of the message
}
