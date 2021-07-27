package packet

import (
	"fmt"
	"sync"
)

// IPacketHandler represents all base functions for handling packet being sent from client
// or from server.
type IPacketHandler interface {
	// Handle handles packets which are being send through packet channel
	Handle(pctChan <-chan IPacket)
}

// IClientHandler contains all methods related to handle client sent packets
type IClientHandler interface {

}

// IServerHandler contains all methods related to handle server sent packets
type IServerHandler interface {

}

// ClientHandler is handler for all packets which comes from client
type ClientHandler struct {
	Wg sync.WaitGroup
}

func (c *ClientHandler) Handle(pctChan <-chan IPacket) {
	defer c.Wg.Done()
	for pct := range pctChan {
		fmt.Printf("Client-> PID: %v, Type: %v, Payload: %#x, String->: %s\n", pct.GetPid(), pct.GetType(), pct.GetPayload(), string(pct.GetPayload()))
	}
}

// ServerHandler is handler for all packets which comes from server and need
// to be send to client
type ServerHandler struct {
	Wg sync.WaitGroup
}

func (s *ServerHandler) Handle(pctChan <-chan IPacket) {
	defer s.Wg.Done()
	for pct := range pctChan {
		fmt.Printf("Server-> PID: %v, Type: %v, Payload: %#x, String->: %s\n", pct.GetPid(), pct.GetType(), pct.GetPayload(), string(pct.GetPayload()))
	}
}
