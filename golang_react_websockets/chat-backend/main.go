package main

import (
	"encoding/json"
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

type Client struct {
	Conn     *websocket.Conn
	Group    string
	Username string
}

var clients = make(map[Client]bool)
var lock = sync.Mutex{}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error upgrading connection:", err)
		return
	}

	group := r.URL.Query().Get("group")
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "anonymous"
	}
	client := Client{Conn: conn, Group: group, Username: username}

	lock.Lock()
	clients[client] = true
	lock.Unlock()

	// Broadcast updated user list to the group
	broadcastUsers(group)

	defer func() {
		lock.Lock()
		delete(clients, client)
		lock.Unlock()
		conn.Close()
		// Broadcast updated user list to the group after user disconnects
		broadcastUsers(group)
	}()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var message map[string]string
		json.Unmarshal(msg, &message)
		switch message["type"] {
		case "message":
			message["content"] = message["content"] + " from " + client.Username
			bresponse, err := json.Marshal(message)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Message received:", message["content"])
			broadcast(messageType, bresponse)
		default:
			broadcastUsers(group)
			log.Println("Unknown message type:", message["type"])
		}
	}
}

func broadcast(messageType int, msg []byte) {
	log.Println("broadcasting message")
	lock.Lock()
	defer lock.Unlock()

	for client := range clients {
		err := client.Conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Printf("Error broadcasting message to client: %v", err)
			client.Conn.Close()
			delete(clients, client)
		}
	}
}

func broadcastUsers(group string) {
	users := getUsersForGroup(group)
	message, _ := json.Marshal(map[string]interface{}{
		"type":  "users_list",
		"users": users,
	})

	lock.Lock()
	defer lock.Unlock()

	for client := range clients {
		if client.Group == group {
			client.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func getUsersForGroup(group string) []string {
	lock.Lock()
	defer lock.Unlock()

	var usernames []string
	for client := range clients {
		if client.Group == group {
			usernames = append(usernames, client.Username)
		}
	}
	return usernames
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
