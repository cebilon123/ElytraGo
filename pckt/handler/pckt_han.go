package handler

import (
	"github.com/cebilon123/ElytraGo/config"
	"github.com/cebilon123/ElytraGo/pckt"
	"log"
	"net"
	"os"
	"strconv"
)

type PacketHandlerFunc func(pct pckt.Packet, c net.Conn)

type GeneralPacketHandler interface {
	HandleWithConnection(PacketHandlerFunc, pckt.Packet, net.Conn)
}

var AvailableHandlers chan GeneralPacketHandler

// PacketHandler responsible for handling packets.
type PacketHandler struct {
	id   int       // id of handler (mostly for debug only)
	done chan bool // done channel to inform for eventually close
}

func (p PacketHandler) HandleWithConnection(f PacketHandlerFunc, pct pckt.Packet, c net.Conn) {
	log.Printf("Handler (id: %d) executing %v", p.id, f)
	f(pct,c)
}

// init is responsible for initialize of handler package.
func init() {
	maxCPH, err := strconv.Atoi(os.Getenv(config.MaxConcurrentPacketHandlersKey))
	if err != nil {
		maxCPH = 50
	}

	// create buffered channel of available handlers to avoid deadlocks.
	AvailableHandlers = make(chan GeneralPacketHandler, maxCPH)
	for i := 0; i < maxCPH; i++ {
		AvailableHandlers <- PacketHandler{
			id:   i,
			done: make(chan bool),
		}
	}
}
