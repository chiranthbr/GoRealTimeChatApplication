package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	CheckOrigin: func (r *http.Request) bool  {
		return true
	},
}

var usernames []string
var clients map[string]*websocket.Conn

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading: ", err)
		return
	}
	defer conn.Close()

	if _, exists := clients["chiru"]; !exists {
		clients["chiru"] = conn
	} else if _, exists := clients["praju"]; !exists {
		clients["praju"] = conn
	}
	

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message! ", err)
			break
		}

		fmt.Println("Recieved message! : ", message)

		if err := conn.WriteMessage(websocket.TextMessage, message[:1]); err != nil {
			fmt.Println("Error sending message to client: ", err)
			break
		}
	}

}

func main() {
	fmt.Println("Server!")
	fmt.Println("Server started at localhost:8080!")
	usernames = append(usernames, "chiru", "praju")

	http.HandleFunc("/chat", wsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server : ", err)
	}
}



