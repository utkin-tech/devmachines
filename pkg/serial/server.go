package serial

import (
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer wsConn.Close()

	unixConn, err := net.Dial("unix", unixSocketPath)
	if err != nil {
		log.Println("Unix socket error:", err)
		wsConn.WriteMessage(websocket.TextMessage, []byte("Failed to connect to Unix socket"))
		return
	}
	defer unixConn.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := unixConn.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Println("Read error:", err)
				}
				wsConn.WriteMessage(websocket.TextMessage, []byte("Socket closed"))
				wsConn.Close()
				return
			}
			wsConn.WriteMessage(websocket.BinaryMessage, buf[:n])
		}
	}()

	for {
		_, msg, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		_, err = unixConn.Write(msg)
		if err != nil {
			log.Println("Write to Unix socket error:", err)
			break
		}
	}
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handleWebSocket)
	mux.Handle("/", http.FileServer(http.Dir("./static/serial")))

	go http.ListenAndServe(":8080", mux)
}
