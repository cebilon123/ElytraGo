package conn

import (
	"bufio"
	"fmt"
	"github.com/cebilon123/ElytraGo/packet"
	"net"
)

type Connection struct {

}

func HandleConnection(c net.Conn) {
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

		pct := packet.NewPacket(pctBytes)

		fmt.Printf("Length: %d | %#v, %s %s, Type: %v\n", length, pctBytes, "string ->", string(pctBytes), pct.GetType())
	}

	c.Close()
}
