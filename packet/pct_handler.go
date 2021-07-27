package packet

import "sync"

// ClSrvChan stores channels
//	PacketsFC is channel for packets which comes from client (needs to be handled by server)
// 	PacketsTC is channel for packets which comes from server (needs to be send to client)
type ClSrvChan struct {
	PacketsFC chan IPacket
	PacketsTC chan IPacket
}

// NewClSrvChan creates new instance of client-server channel handler
func NewClSrvChan() *ClSrvChan {
	return &ClSrvChan{
		PacketsFC: make(chan IPacket, 1),
		PacketsTC: make(chan IPacket, 1),
	}
}

// Close closes channels
func (c ClSrvChan) Close() error {
	close(c.PacketsFC)
	close(c.PacketsTC)

	return nil
}

// StartPacketListeningAndHandling registers handlers for channel communication.
// 	See code for information which handlers are being used.
func (c ClSrvChan) StartPacketListeningAndHandling() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	cHan := &ClientHandler{&wg}
	sHan := &ServerHandler{&wg}

	select {
	case <-c.PacketsFC:
		cHan.Handle(c.PacketsFC)
	case <-c.PacketsTC:
		sHan.Handle(c.PacketsTC)
	}

	wg.Wait()
}
