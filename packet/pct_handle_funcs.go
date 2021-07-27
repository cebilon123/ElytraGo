package packet

// Handle is delegate to func which handles pct. Second parameter is chan which handles
// responses
type Handle func(pct IPacket, sChan <-chan IPacket)

// TypeHandleFuncMap stores packetType:handleFunc
var TypeHandleFuncMap = map[Type]Handle{}
