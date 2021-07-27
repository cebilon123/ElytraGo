package packet

import "sync"

// ClSrvChan stores channels
//	PacketsFC is channel for packets which comes from client (needs to be handled by server)
// 	PacketsFS is channel for packets which comes from server (needs to be send to client)
type ClSrvChan struct {
	PacketsFC chan IPacket
	PacketsFS chan IPacket
}

// NewClSrvChan creates new instance of client-server channel handler
func NewClSrvChan() *ClSrvChan {
	return &ClSrvChan{
		PacketsFC: make(chan IPacket, 1),
		PacketsFS: make(chan IPacket, 1),
	}
}

// Close closes channels
func (c ClSrvChan) Close() error {
	close(c.PacketsFC)
	close(c.PacketsFS)

	return nil
}

// HandleChannelsCommunication registers handlers for channel communication.
// 	See code for information which handlers are being used.
func (c ClSrvChan) HandleChannelsCommunication() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	cHan := &ClientHandler{}
	sHan := &ServerHandler{}

	select {
	case <-c.PacketsFC:
		cHan.Handle(c.PacketsFC)
	case <-c.PacketsFS:
		sHan.Handle(c.PacketsFS)
	}

	wg.Wait()
}
