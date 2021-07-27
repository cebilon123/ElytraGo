package packet

// Handle is delegate to func which handles pct. Second parameter is chan which handles
// responses
type Handle func(pct IPacket, sChan chan<- IPacket, c *ClientHandler) error

// TypeHandleFuncMap stores packetType:handleFunc
var TypeHandleFuncMap = map[Type]Handle{
	ServerStatusC: ServerStatusCHandle,
}


// ServerStatusCHandle returns server status to client
func ServerStatusCHandle(pct IPacket, sChan chan<- IPacket, c *ClientHandler) error {


	return nil
}