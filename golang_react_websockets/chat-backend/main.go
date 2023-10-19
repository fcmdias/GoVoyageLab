package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var lock = sync.Mutex{}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error upgrading connection:", err)
		return
	}
	defer func() {
		lock.Lock()
		delete(clients, conn)
		lock.Unlock()
		conn.Close()
	}()

	lock.Lock()
	clients[conn] = true
	lock.Unlock()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		broadcast(messageType, msg)
	}
}

func broadcast(messageType int, msg []byte) {
	lock.Lock()
	defer lock.Unlock()

	for client := range clients {
		err := client.WriteMessage(messageType, msg)
		if err != nil {
			log.Printf("Error broadcasting message to client: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
