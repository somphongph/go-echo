package publicapis

type rootResponse struct {
	Count   int             `json:"count"`
	Entries []entryResponse `json:"entries"`
}

type entryResponse struct {
	API         string `json:"api"`
	Description string `json:"description"`
	Auth        string `json:"auth"`
	HTTPS       bool   `json:"https"`
	Cors        string `json:"cors"`
	Link        string `json:"link"`
	Category    string `json:"category"`
}
