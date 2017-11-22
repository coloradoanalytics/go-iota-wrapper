package iota

import (
	"github.com/willf/pad"
)

type Transactions struct {
	Hashes []string `json:"hashes"`
}

func (c *Client) FindTransactions(bundles, addresses, tags, approvees []string) (Transactions, error) {

	for i, b := range bundles {
		bundles[i] = pad.Right(b, 81, "9")
	}

	command := Command{
		Command:   "getTransactions",
		Bundles:   bundles,
		Addresses: addresses,
		Tags:      tags,
		Approvees: approvees,
	}

	transactions := Transactions{}

	err := c.do(command, &transactions)

	return transactions, err
}

type TransactionsToApprove struct {
	TrunkTransaction  string `json:"trunkTransaction"`
	BranchTransaction string `json:"branchTransaction"`
	Duration          int    `json:"duration"`
}

func (c *Client) GetTransactionsToApprove(depth int) (TransactionsToApprove, error) {
	command := Command{
		Command: "getTransactionsToApprove",
		Depth:   depth,
	}

	transactions := TransactionsToApprove{}

	err := c.do(command, &transactions)

	return transactions, err
}

func (c *Client) BroadcastTransactions(trytes []string) error {
	command := Command{
		Command: "broadcastTransactions",
		Trytes:  trytes,
	}

	err := c.do(command, nil)

	return err
}

func (c *Client) StoreTransactions(trytes []string) error {
	command := Command{
		Command: "storeTransactions",
		Trytes:  trytes,
	}

	err := c.do(command, nil)

	return err
}
