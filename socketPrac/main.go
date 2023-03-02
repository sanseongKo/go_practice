package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Map struct {
	wsMap map[Ws]bool
}

type Ws struct {
	ws *websocket.Conn
}

var singleWsInstance = &Map{
	wsMap: make(map[Ws]bool),
}
var lock = &sync.Mutex{}

func getInstance(ws *websocket.Conn) *Map {
	lock.Lock()
	defer lock.Unlock()

	if singleWsInstance == nil {
		singleWsInstance.wsMap[Ws{ws: ws}] = true
		return singleWsInstance
	} else {
		return singleWsInstance
	}
}

func main() {
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc("/make", something)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// upgrade this connection to a WebSocket
	// connection
	ws, _ := upgrader.Upgrade(w, r, nil)
	testMap := getInstance(ws)
	wss := Ws{ws}
	testMap.wsMap[wss] = true

	var text string
	defer r.Body.Close()

	a, err1 := io.ReadAll(r.Body)

	if err1 != nil {

	}

	json.Unmarshal([]byte(a), &text)

	if text == "" {
		text = "no text"
	}

	log.Println("Client Connected")

	writeMessage(ws, text)
}

func writeMessage(ws *websocket.Conn, text string) {
	err := ws.WriteMessage(1, []byte("Hello client"+text+"!"))
	if err != nil {
		log.Println(err)
	}
	reader(ws)

	defer ws.Close()
}

func something(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer r.Body.Close()
	testMap := getInstance(ws)
	a, err1 := io.ReadAll(r.Body)
	text := bytes.NewBuffer(a).String()

	if err1 != nil {

	}

	if text == "" {
		text = "no text"
	}

	for key, val := range testMap.wsMap {
		if val == true {
			go writeMessage(key.ws, text)
		}
	}

	log.Println("Client Connected")

	res := text
	json.NewEncoder(w).Encode(res)
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out incoming message
		fmt.Println("test incoming message: " + string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
