package client

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	imgColor "image/color"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id    string
	Color string
	Hub   *Hub
	conn  *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Setup() {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		color.Red("Error upgrading connection: %s", err)
		return
	}
	colors := []imgColor.Color{imgColor.RGBA{255, 0, 0, 255}, imgColor.RGBA{0, 255, 0, 255}, imgColor.RGBA{0, 0, 255, 255}, imgColor.RGBA{255, 255, 0, 255}, imgColor.RGBA{255, 0, 255, 255}, imgColor.RGBA{0, 255, 255, 255}}

	// create a new client
	client := &Client{
		Id:    uuid.New().String(),
		Color: toHex(colors[rand.Intn(len(colors))]),
		conn:  conn,
		Hub:   hub,
	}

	// register the client with the hub
	hub.register <- client

	// send all previous messages to the client
	for _, message := range hub.Messages {
		err := conn.WriteJSON(message)
		if err != nil {
			color.Red("Error sending message: %s", err)
			break
		}
	}

	// listen for messages from the client
	go client.Listen()
}

func toHex(c imgColor.Color) string {
	r, g, b, _ := c.RGBA()
	r = r / 256
	g = g / 256
	b = b / 256
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func (c *Client) Listen() {
	defer func() {
		c.Hub.unregister <- c
		c.conn.Close()
	}()

	for {
		message := Message{}
		err := c.conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived) {
				color.Red("Error reading message: %s", err)
			}
			break
		}

		// set the Message Created field
		message.Created = time.Now().Format("2006-01-02 15:04:05")

		//set the color of the message to be a random color
		message.Color = c.Color

		// broadcast the message to all clients
		c.Hub.broadcast <- message

		// add the message to the hub
		c.Hub.Messages = append(c.Hub.Messages, message)
	}
}
