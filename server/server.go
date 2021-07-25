package server

import (
	"github.com/cebilon123/ElytraGo/conn"
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
	l, err := net.Listen("tcp", ":9999")

	if err != nil {
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return
		}
		go conn.HandleConnection(c)
	}
}


