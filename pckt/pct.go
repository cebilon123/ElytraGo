package pckt

// Packet represents packet which is send from server to client
// and from client to server
type Packet struct {
	Len        int64
	Id         int64
	Data       []byte
	FromClient bool
}

// New creates new packet
func New(len int64, id int64, data []byte, fromClient bool) *Packet {
	return &Packet{Len: len, Id: id, Data: data, FromClient: fromClient}
}

