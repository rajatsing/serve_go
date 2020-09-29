package main

import (
	"net/http"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func webSocket(res http.ResponseWriter, req *http.Request) {
	// Gorilla Websocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, error := upgrader.Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)                // namegenerator API call
	name := nameGenerator.Generate()                                     // Generating a random Name
	client := &Client{name: name, socket: conn, send: make(chan []byte)} // client is created with a unique name
	manager.register <- client                                           // registering the newly created client

	go client.read()  // read the message from the client
	go client.write() // write the message from client to the dashboard

}
