// Code generated by goctl. DO NOT EDIT.
package types

type TransactionArgs struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	Gas                  uint64 `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	Value                string `json:"value"`
	Nonce                uint64 `json:"nonce"`
	SigAlgo              byte   `json:"sigAlgo"`
	Signature            string `json:"signature"`
	Data                 string `json:"data"`
	Input                string `json:"input"`
	AccessList           string `json:"accessList,omitemty"`
	ChainID              string `json:"chainId,omitempty"`
}

type BoolRes struct {
	Flag bool `json:"flag"`
}
