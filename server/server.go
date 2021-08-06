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
	defer l.Close()
	if err != nil {
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			return
		}
		go conn.HandleConnection(c)
	}
}
