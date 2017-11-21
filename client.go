package iota

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	URI        string //API address of IRI, including port
	httpClient *http.Client
}

func NewClient(uri string) *Client {
	return &Client{
		URI:        uri,
		httpClient: &http.Client{},
	}
}

type Command struct {
	Addresses          []string `json:"addresses,omitempty"`
	Approvees          []string `json:"approvees,omitempty"`
	BranchTransaction  string   `json:"branchTransaction,omitempty"`
	Bundles            []string `json:"bundles,omitempty"`
	Command            string   `json:"command"`
	Depth              int      `json:"depth,omitempty"`
	Hashes             []string `json:"hashes,omitempty"`
	MinWeightMagnitude int      `json:"minWeightMagnitude,omitempty"`
	Tags               []string `json:"tags,omitempty"`
	Threshold          int      `json:"threshold,omitempty"`
	Tips               []string `json:"tips,omitempty"`
	Transactions       []string `json:"transactions,omitempty"`
	TrunkTransaction   string   `json:"trunkTransaction,omitempty"`
	Trytes             []string `json:"trytes,omitempty"`
	URIs               []string `json:"uris,omitempty"`
}

func (c *Client) do(command Command, result interface{}) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(command)

	req, err := http.NewRequest("POST", c.URI, b)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-IOTA-API-Version", "1")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	json.NewDecoder(res.Body).Decode(result)

	return nil
}
