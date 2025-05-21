package vnc

import (
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer wsConn.Close()

	unixConn, err := net.Dial("unix", unixSocketPath)
	if err != nil {
		log.Println("Unix socket connection error:", err)
		return
	}
	defer unixConn.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := unixConn.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Println("Unix socket read error:", err)
				}
				return
			}
			err = wsConn.WriteMessage(websocket.BinaryMessage, buf[:n])
			if err != nil {
				log.Println("WebSocket write error:", err)
				return
			}
		}
	}()

	for {
		_, msg, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			unixConn.Close()
			return
		}
		_, err = unixConn.Write(msg)
		if err != nil {
			log.Println("Unix socket write error:", err)
			return
		}
	}
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/websockify", proxyHandler)
	mux.Handle("/", http.FileServer(http.Dir("./static/vnc")))

	go http.ListenAndServe(":8081", mux)
}
