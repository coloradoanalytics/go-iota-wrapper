package iota

type Balances struct {
	Balances       []string `json:"balances"`
	Duration       int      `json:"duration"`
	Milestone      string   `json:"milestone"`
	MilestoneIndex int      `json:"milestoneIndex"`
}

func (c *Client) GetBalances(addresses []string) (Balances, error) {
	command := Command{
		Command:   "getBalances",
		Addresses: addresses,
		Threshold: 100,
	}

	balances := Balances{}

	err := c.do(command, &balances)

	return balances, err
}
