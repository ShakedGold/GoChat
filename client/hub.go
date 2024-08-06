package client

type Hub struct {
	Messages   []Message
	Clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

type Message struct {
	From    string `json:"from"`
	Content string `json:"content"`
	Created string `json:"created"`
	Color   string `json:"color"`
}

func NewHub() *Hub {
	return &Hub{
		Messages:   []Message{},
		broadcast:  make(chan Message),
		Clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Clients[client] = true
		case message := <-h.broadcast:
			for client := range h.Clients {
				client.conn.WriteJSON(message)
			}
		}
	}
}
