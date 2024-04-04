package main

import "net/http"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveHome(w http.ResponseWriter, r *http.Request) {

}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil){
		if err != nil{
			// Handle error
		}
	}
	defer ws.Close()
}
