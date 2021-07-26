package packet

// IPacket specifies signatures for packets of various kind.
// Starting of handshake etc.
type IPacket interface {
	GetPayload() []byte
	GetPid() byte
	GetType() ClientPacket
}

// Packet its base struct which represents Packet.
//  ClientPacket represents enum which is just converted PId
// 	PId represents packetId
//	Payload represents Payload which is being send through Packet
type Packet struct {
	Type    ClientPacket
	PId     byte
	Payload []byte
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
func (p Packet) GetType() ClientPacket {
	return p.Type
}

// ClientPacket is enum
type ClientPacket byte

const (
	Default ClientPacket = iota
	Handshake
)

// NewPacket creates Packet, which type is being resolved based on Packet []byte.
// This func also sets ClientPacket as Type in IPacket, as well its sets payload
// (which ignores first byte and sets it as it should be set. Payload is still
// []byte which needs to be converted in later usages of value)
func NewPacket(packet []byte) IPacket {
	var resolvePayload = func() []byte {
		if len(packet) < 2 {
			return make([]byte, 0)
		}

		return packet[1:]
	}

	return &Packet{
		Type:    resolveClientPacket(packet),
		PId:     packet[0],
		Payload: resolvePayload(),
	}
}

// resolveClientPacket returns ClientPacket based on the first byte
// of Packet being send as parameter.
func resolveClientPacket(packet []byte) ClientPacket {
	switch packet[0] {
	case 0x0:
		return Handshake

	}

	return Default
}
