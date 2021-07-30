package packet

import "net"

// IPacket specifies signatures for packets of various kind.
// Starting of handshake etc.
type IPacket interface {
	GetPayload() []byte
	GetPid() byte
	GetType() Type
	GetConn() net.Conn
}

// Packet its base struct which represents Packet.
//  Type represents enum which is just converted PId
// 	PId represents packetId
//	Payload represents Payload which is being send through Packet
type Packet struct {
	Type    Type
	PId     byte
	Payload []byte
	c       net.Conn
}

// GetPayload returns packet payload.
func (p Packet) GetPayload() []byte {
	return p.Payload
}

// GetPid returns PId of the packet
func (p Packet) GetPid() byte {
	return p.PId
}

// GetType returns type of the packet
func (p Packet) GetType() Type {
	return p.Type
}

// GetConn returns current connection
func (p Packet) GetConn() net.Conn {
	return p.c
}

// Type is enum
type Type byte

// Suffix "C" or "S" represents if type comes from client or from server
const (
	Default Type = iota
	HandshakeC
	HandshakeS
	ServerStatusC
)

// NewPacket creates Packet, which type is being resolved based on Packet []byte.
// This func also sets Type as Type in IPacket, as well its sets payload
// (which ignores first byte and sets it as it should be set. Payload is still
// []byte which needs to be converted in later usages of value)
//
//	IMPORTANT: clientPct need to be set in case of correct choose of type of the packet
func NewPacket(packet []byte, clientPct bool, c net.Conn) IPacket {
	var resolvePayload = func() []byte {
		if len(packet) < 2 {
			return make([]byte, 0)
		}

		return packet[1:]
	}

	var resolvePctType = func() Type {
		if clientPct {
			return resolveClientPacket(packet)
		}
		return resolveServerPacket(packet)
	}

	return &Packet{
		Type:    resolvePctType(),
		PId:     packet[0],
		Payload: resolvePayload(),
		c: c,
	}
}

// resolveClientPacket returns Type based on the first byte
// of Packet being send as parameter.
func resolveClientPacket(packet []byte) Type {
	switch packet[0] {
	case 0x0:
		if len(packet) > 5 {
			return HandshakeC
		}

		return ServerStatusC
	}

	return Default
}

// resolveServerPacket returns Type based on the first byte
// of Packet being send as parameter.
func resolveServerPacket(packet []byte) Type {
	switch packet[0] {
	case 0x0:
		return HandshakeS // or just empty message but currently for the sake of simplicity we just have this HandshakeS
	}

	return Default
}
