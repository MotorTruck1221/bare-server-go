package v3

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Handshake struct {
	Type           string            `json:"type"`
	Remote         string            `json:"remote"`
	Protocols      []string          `json:"protocols"`
	Headers        map[string]string `json:"headers"`
	ForwardHeaders []string          `json:"forwardHeaders"`
}

type Response struct {
	Type       string   `json:"type"`
	Protocol   string   `json:"protocol"`
	SetCookies []string `json:"setCookies"`
}

func forwardMessages(source *websocket.Conn, destination *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		messageType, message, err := source.ReadMessage()
		if err != nil {
			fmt.Println("Error Reading Message", err)
			break
		}
		err = destination.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Error Writing Message", err)
			break
		}
	}
}

func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	// upgrade to websocket
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	//read from the client
	clientHandshake := Handshake{}
	err = conn.ReadJSON(&clientHandshake)
	// fmt.Println(clientHandshake.Headers)
	if err != nil {
		panic(err)
	}
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 0,
	}

	headers := make(http.Header)
	for key, value := range clientHandshake.Headers {
		if key != "Upgrade" && key != "Connection" {
			headers.Set(key, value)
		}
	}
	// fmt.Println(headers)

	remoteConn, resp, err := dialer.Dial(
		clientHandshake.Remote,
		headers,
	)
	if err != nil {
        if err == websocket.ErrBadHandshake {
            fmt.Println("Bad Handshake", resp.StatusCode)
            fmt.Println("Response Status", resp.Status)
            fmt.Println(headers)
        }
	}
	fmt.Println("Connected to remote server", remoteConn.RemoteAddr())
	fmt.Println("Sending response to client")
	response := Response{
		Type:       "open",
		Protocol:   "",
		SetCookies: []string{},
	}
	err = conn.WriteJSON(response)
	if err != nil {
		panic(err)
	}

	//start forwarding messages
	var wg sync.WaitGroup
	wg.Add(2)
	go forwardMessages(conn, remoteConn, &wg)
	go forwardMessages(remoteConn, conn, &wg)
	wg.Wait()
	//defer conn.Close()
}
