package fireblocks

type CreateTransactionRequest struct {
	Operation          string         `json:"operation"`
	Note               string         `json:"note"`
	ExternalTxID       string         `json:"externalTxId"`
	AssetID            string         `json:"assetId"`
	Source             Source         `json:"source"`
	Destination        Destination    `json:"destination"`
	Destinations       []Destinations `json:"destinations"`
	Amount             string         `json:"amount"`
	TreatAsGrossAmount bool           `json:"treatAsGrossAmount"`
	ForceSweep         bool           `json:"forceSweep"`
	FeeLevel           string         `json:"feeLevel"`
	Fee                string         `json:"fee"`
	PriorityFee        string         `json:"priorityFee"`
	FailOnLowFee       bool           `json:"failOnLowFee"`
	MaxFee             string         `json:"maxFee"`
	GasLimit           string         `json:"gasLimit"`
	GasPrice           string         `json:"gasPrice"`
	NetworkFee         string         `json:"networkFee"`
	ReplaceTxByHash    string         `json:"replaceTxByHash"`
	CustomerRefID      string         `json:"customerRefId"`
	ExtraParameters    RawMessageData `json:"extraParameters"`
}

type Source struct {
	Type     string `json:"type"`
	SubType  string `json:"subType"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	WalletID string `json:"walletId"`
}

type Destination struct {
	Type           string         `json:"type"`
	SubType        string         `json:"subType"`
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	WalletID       string         `json:"walletId"`
	OneTimeAddress OneTimeAddress `json:"oneTimeAddress"`
}

type Destinations struct {
	Amount      string      `json:"amount"`
	Destination Destination `json:"destination"`
}

type OneTimeAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

type RawMessageData struct {
	Messages  []UnsignedRawMessage `json:"messages"`
	Algorithm string               `json:"algorithm"`
}

type UnsignedRawMessage struct {
	Content           string    `json:"content"`
	Bip44addressIndex float64   `json:"bip44AddressIndex"`
	Bip44change       float64   `json:"bip44Change"`
	DerivationPath    []float64 `json:"derivationPath"`
}

type TransactionResponse struct {
	ID             string         `json:"id"`
	Status         string         `json:"status"`
	SystemMessages SystemMessages `json:"systemMessages"`
}
type SystemMessages struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
