package goiw

import ()

type Tips struct {
	Hashes   []string `json:"hashes"`
	Duration int      `json:"duration"`
}

func (c *Client) GetTips() (Tips, error) {
	command := Command{
		Command: "getTips",
	}

	tips := Tips{}

	err := c.do(command, &tips)

	return tips, err
}
