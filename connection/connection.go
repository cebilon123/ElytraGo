package connection

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConnection(c net.Conn) {
	fmt.Printf("Serving: %s\n", c.RemoteAddr().String())

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", netData)

		temp := strings.TrimSpace(netData)
		if temp == "STOP" {
			break
		}

		c.Write([]byte(netData))
	}

	c.Close()
}
