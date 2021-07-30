package packet

func CreateServerStatusResponse() *ServerStatusResponse {
	return &ServerStatusResponse{
		Version: struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		}{
			Name:     "1.16.1",
			Protocol: 48,
		},
		Players: struct {
			Max    int `json:"max"`
			Online int `json:"online"`
			Sample []struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"sample"`
		}{
			Max:    24,
			Online: 3,
			Sample: []struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			}{
				{
					Name: "test",
					Id:   "4566e69f-c907-48ee-8d71-d7ba5aa00d20",
				},
			},
		},
		Description: struct {
			Text string `json:"text"`
		}{
			Text: "Hello world from ElytraGO",
		},
		Favicon: "",
	}
}
