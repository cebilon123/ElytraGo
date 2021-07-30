package server

import (
	"github.com/cebilon123/ElytraGo/conn"
	"github.com/cebilon123/ElytraGo/packet"
	"net"
)

// IBaseServerBuilder represents functionality for building server in builder manner.
type IBaseServerBuilder interface {
	Create() IBaseServerBuilder
	Start()
}

type Builder struct {

}

func NewBuilder() *Builder {
	return &Builder{}
}

func (s Builder) Create() IBaseServerBuilder {
	return s
}

func (s Builder) Start() {
	wd := conn.NewWorkerDispatcher(make(chan packet.IPacket), make(chan packet.IPacket))
	go wd.SpawnWorkers()
	defer wd.Close()

	l, err := net.Listen("tcp", ":9999")
	defer l.Close()
	if err != nil {
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			return
		}
		go conn.HandleConnection(c, wd)
	}
}


