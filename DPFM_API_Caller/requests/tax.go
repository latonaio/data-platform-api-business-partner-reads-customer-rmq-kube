package requests

type Tax struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	Customer            int     `json:"Customer"`
	DepartureCountry    string  `json:"DepartureCountry"`
	TaxCategory         *string `json:"TaxCategory"`
	BPTaxClassification *string `json:"BPTaxClassification"`
}
