package iota

type Trytes struct {
	Trytes []string `json:"trytes"`
}

func (c *Client) GetTrytes(hashes []string) (Trytes, error) {
	command := Command{
		Command: "getTrytes",
		Hashes:  hashes,
	}

	trytes := Trytes{}

	err := c.do(command, &trytes)

	return trytes, err
}
