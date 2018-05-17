package volleynet

type Client struct {
	PostUrl     string
	ApiUrl      string
	AmateurPath string
	Cookie      string
}

func DefaultClient() *Client {
	return &Client{
		PostUrl:     "https://beach.volleynet.at/Admin/formular",
		ApiUrl:      "http://www.volleynet.at/api/",
		AmateurPath: "beach/bewerbe/%s/phase/%s/sex/%s/saison/%s/information/all",
	}
}
