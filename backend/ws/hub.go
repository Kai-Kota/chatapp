package ws

type Hub struct {
	clients    map[*MessageClient]bool
	broadcast  chan []byte
	register   chan *MessageClient
	unregister chan *MessageClient
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *MessageClient),
		unregister: make(chan *MessageClient),
		clients:    make(map[*MessageClient]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
