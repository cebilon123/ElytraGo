package conn

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/cebilon123/ElytraGo/pckt"
	"github.com/cebilon123/ElytraGo/pckt/handler"
	"io/ioutil"
	"log"
	"net"
)

// HandleConnection handles connection client->server and vice versa.
func HandleConnection(c net.Conn) {
	defer closeConn(c)
	fmt.Printf("\nServing: %s\n", c.RemoteAddr().String())

	for {
		pctBytes := readBytesFromConnection(c) // read all byes from connection

		pctLen, readIndex := binary.Varint(pctBytes)              // pull var int to get length of pctId+data
		pctId, readIndexId := binary.Varint(pctBytes[readIndex:]) // next pull pctId
		pctData := pctBytes[(readIndexId + readIndex):]           // rest of bytes are payload

		pct := pckt.New(pctLen, pctId, pctData, true)
		h := <-handler.AvailableHandlers // pull any available handler from channel
		go func() {
			funcForPct := getHandlerForPacket(*pct) // choose handler func based on packet id and data
			h.HandleWithConnection(funcForPct)      // handle this function
			handler.AvailableHandlers <- h          // when all work is done we are putting handler back on queue so this way we have always equal amount of handlers
		}()
	}
}

// closeConn closes current connection and prints errors if any appears.
func closeConn(c net.Conn) {
	func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}(c)
}

// readBytesFromConnection reads all bytes from connection and return them in bytes slice.
func readBytesFromConnection(c net.Conn) []byte {
	reader := bufio.NewReader(c)
	pctBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return pctBytes
}

// getHandlerForPacket chooses right handler for any packet. TODO: there is need to create handlers repository and code to choose right one
func getHandlerForPacket(pct pckt.Packet) handler.PacketHandlerFunc {
	return func(pct pckt.Packet, c net.Conn) {

	}
}
