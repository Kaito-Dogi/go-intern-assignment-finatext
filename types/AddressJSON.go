package types

type AddressJSON struct {
	PostalCode  string `json:"postal_code"`
	Address     string `json:"address"`
	AddressKana string `json:"address_kana"`
}
