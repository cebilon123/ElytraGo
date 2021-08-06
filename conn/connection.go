package conn

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

// HandleConnection handles connection client->server and vice versa
func HandleConnection(c net.Conn, wd *WorkerDispatcher) {
	defer closeConn(c)
	fmt.Printf("\nServing: %s\n", c.RemoteAddr().String())

	for {
		pctBytes := readBytesFromConnection(c) // read all byes from connection

		pctLen, readIndex := binary.Varint(pctBytes) // pull var int to get length of pctId+data
		pctId, readIndexId := binary.Varint(pctBytes[readIndex:]) // next pull pctId
		pctData := pctBytes[(readIndexId + readIndex):] // rest of bytes are payload

		fmt.Printf("\n%d, %d\n", pctLen, pctId, pctData)
	}
}

func closeConn(c net.Conn) {
	func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}(c)
}

func readBytesFromConnection(c net.Conn) []byte {
	reader := bufio.NewReader(c)
	pctBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return pctBytes
}
