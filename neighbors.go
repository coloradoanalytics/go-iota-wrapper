package goiw

import ()

type Neighbor struct {
	Address                           string `json:"address"`
	ConnectionType                    string `json:"connectionType"`
	NumberOfAllTransactions           int    `json:"numberOfAllTransactions"`
	NumberOfInvalidTransactions       int    `json:"numberOfInvalidTransactions"`
	NumberOfNewTransactions           int    `json:"numberOfNewTransactions"`
	NumberOfRandomTransactionRequests int    `json:"numberOfRandomTransactionRequests"`
	NumberOfSentTransactions          int    `json:"numberOfSentTransactions"`
}

type Neighbors struct {
	Neighbors []Neighbor `json:"neighbors"`
	Duration  int        `json:"duration"`
}

func (c *Client) GetNeighbors() (Neighbors, error) {
	command := Command{
		Command: "getNeighbors",
	}

	neighbors := Neighbors{}
	err := c.do(command, &neighbors)

	return neighbors, err
}

func (c *Client) AddNeighbors(uris []string) error {
	return nil
}

func (c *Client) RemoveNeighbors(uris []string) error {
	return nil
}
