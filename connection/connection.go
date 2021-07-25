package connection

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
}

var tmpJson = "{\n    \"version\": {\n        \"name\": \"1.8.7\",\n        \"protocol\": 47\n    },\n    \"players\": {\n        \"max\": 100,\n        \"online\": 5,\n        \"sample\": [\n            {\n                \"name\": \"thinkofdeath\",\n                \"id\": \"4566e69f-c907-48ee-8d71-d7ba5aa00d20\"\n            }\n        ]\n    },\n    \"description\": {\n        \"text\": \"Hello world\"\n    },\n    \"favicon\": \"data:image/png;base64,<data>\"\n}"

type Status struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
		Sample []struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"sample"`
	} `json:"players"`
	Description struct {
		Text string `json:"text"`
	} `json:"description"`
	Favicon string `json:"favicon"`
}

func HandleConnection(c net.Conn) {
	//fmt.Printf("\nServing: %s\n", c.RemoteAddr().String())

	for {
		res := make([]byte, 32)
		_, err := bufio.NewReader(c).Read(res)
		if err != nil {
			fmt.Println(err)
			return
		}

		if res[1] == 0x00 {
			fmt.Printf("\n%#v", res)
			fmt.Printf(" (string) %s\n", res)
		}
	}

	c.Close()
}
