package server

import (
	"github.com/cebilon123/ElytraGo/connection"
	"math/rand"
	"net"
	"time"
)

// IBaseServerBuilder represents functionality for building server in builder manner.
type IBaseServerBuilder interface {
	Create() IBaseServerBuilder
	Start() IBaseServerBuilder
}

type Builder struct {

}

func New() *Builder {
	return &Builder{}
}

func (s Builder) Create() IBaseServerBuilder {
	return s
}

func (s Builder) Start() IBaseServerBuilder {
	l, err := net.Listen("tcp4", ":8084")
	if err != nil {
		return nil
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			return nil
		}
		go connection.HandleConnection(c)
	}

	return s
}


