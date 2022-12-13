package requests

type PartnerPlant struct {
	BusinessPartner                int     `json:"BusinessPartner"`
	Customer                       int     `json:"Customer"`
	PartnerCounter                 int     `json:"PartnerCounter"`
	PartnerFunction                string  `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner int     `json:"PartnerFunctionBusinessPartner"`
	PlantCounter                   int     `json:"PlantCounter"`
	Plant                          *string `json:"Plant"`
	DefaultPlant                   *bool   `json:"DefaultPlant"`
	DefaultStockConfirmationPlant  *bool   `json:"DefaultStockConfirmationPlant"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}
