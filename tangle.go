package goiw

import ()

//see trytes.go for Trytes struct

func (c *Client) AttachToTangle(trunkTransaction string, branchTransaction string, minWeightMagnitude int, trytes []string) (Trytes, error) {

	command := Command{
		Command:            "attachToTangle",
		TrunkTransaction:   trunkTransaction,
		BranchTransaction:  branchTransaction,
		MinWeightMagnitude: minWeightMagnitude,
		Trytes:             trytes,
	}

	t := Trytes{}

	err := c.do(command, &t)

	return t, err
}

//interruptAttachingToTangle has incomplete documentation
