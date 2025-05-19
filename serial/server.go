package serial

import (
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	usock, err := net.Dial("unix", unixSocketPath)
	if err != nil {
		log.Println("Unix socket error:", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Failed to connect to Unix socket"))
		return
	}
	defer usock.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := usock.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Println("Read error:", err)
				}
				conn.WriteMessage(websocket.TextMessage, []byte("Socket closed"))
				conn.Close()
				return
			}
			conn.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		_, err = usock.Write(msg)
		if err != nil {
			log.Println("Write to Unix socket error:", err)
			break
		}
	}
}

func startServer() {
	http.HandleFunc("/ws", handleWebSocket)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	go http.ListenAndServe(":8080", nil)
}
