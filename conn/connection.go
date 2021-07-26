package conn

import (
	"bufio"
	"fmt"
	"github.com/cebilon123/ElytraGo/packet"
	"net"
)

// Connection stores all information about current connection.
//	PacketsFC is channel for packets which comes from client (needs to be handled by server)
// 	PacketsFS is channel for packets which comes from server (needs to be send to client)
type Connection struct {
	PacketsFC chan packet.IPacket
	PacketsFS chan packet.IPacket
}

// HandleConnection handles connection client->server and vice versa
func HandleConnection(c net.Conn) {
	defer c.Close()
	fmt.Printf("\nServing: %s\n", c.RemoteAddr().String())

	for {
		length := make([]byte, 1)
		reader := bufio.NewReader(c)
		_, err := reader.Read(length)
		if err != nil {
			fmt.Println(err)
			return
		}

		pctBytes := make([]byte, length[0])
		_, err = reader.Read(pctBytes)
		if err != nil {
			fmt.Println(err)
			return
		}

		pct := packet.NewPacket(pctBytes, true)

		fmt.Printf("Length: %d | %#v, %s %s, Type: %v\n", length, pctBytes, "string ->", string(pctBytes), pct.GetType())
	}
}
