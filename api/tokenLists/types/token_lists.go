package types

type TokenList struct {
	Name      string   `json:"name"`
	LogoURI   string   `json:"logoURI"`
	Keywords  []string `json:"keywords"`
	Timestamp string   `json:"timestamp"`
	Tokens    []Token  `json:"tokens"`
	Version   Version  `json:"version"`
}

type Token struct {
	ChainID  int    `json:"chainId"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}
