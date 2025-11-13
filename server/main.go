package main

import (
	"encoding/json"
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
var clients = make(map[string]*websocket.Conn)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading: ", err)
		return
	}
	defer conn.Close()

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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var cred struct {
		username string `json:"username"`
		passowrd string `json:"password"`
	}

	if r.Method != http.MethodPost {
		http.Error("Not suitable method request!", http.StatusMethodNotAllowed)
		return
	}	

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

}

func main() {
	fmt.Println("Server!")
	fmt.Println("Server started at localhost:8080!")
	usernames = append(usernames, "chiru", "praju")

	http.HandleFunc("/chat", wsHandler)
	http.HandleFunc("/loginOrSignup", loginHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server : ", err)
	}
}



