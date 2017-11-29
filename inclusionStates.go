package goiw

type InclusionStates struct {
	States   []bool `json:"states"`
	Duration int    `json:"duration"`
}

func (c *Client) GetInclusionStates(transactions, tips []string) (InclusionStates, error) {
	command := Command{
		Command:      "getInclusionStates",
		Transactions: transactions,
		Tips:         tips,
	}

	states := InclusionStates{}

	err := c.do(command, &states)

	return states, err
}
