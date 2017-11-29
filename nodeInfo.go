package goiw

import ()

type NodeInfo struct {
	AppName                            string `json:"appName"`
	AppVersion                         string `json:"appVersion"`
	Duration                           int    `json:"duration"`
	JREAvailableProcessors             int    `json:"jreAvailableProcessors"`
	JREFreeMemory                      int    `json:"jreFreeMemory"`
	JREMaxMemory                       int    `json:"jreMaxMemory"`
	JRETotalMemory                     int    `json:"jreTotalMemory"`
	LatestMilestone                    string `json:"latestMilestone"`
	LatestMilestoneIndex               int    `json:"latestMilestoneIndex"`
	LatestSolidSubtangleMilestone      string `json:"latestSolidSubtangleMilestone"`
	LatestSolidSubtangleMilestoneIndex int    `json:"latestSolidSubtangleMilestoneIndex"`
	Neighbors                          int    `json:"neighbors"`
	PacketsQueueSize                   int    `json:"packetsQueueSize"`
	Time                               int    `json:"time"`
	Tips                               int    `json:"tips"`
	TransactionsToRequest              int    `json:"transactionsToRequest"`
}

func (c *Client) GetNodeInfo() (NodeInfo, error) {
	command := Command{
		Command: "getNodeInfo",
	}

	nodeInfo := NodeInfo{}

	err := c.do(command, &nodeInfo)

	return nodeInfo, err
}
