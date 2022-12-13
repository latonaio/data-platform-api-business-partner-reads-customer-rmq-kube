package requests

type Accounting struct {
	BusinessPartner       int     `json:"BusinessPartner"`
	Customer              int     `json:"Customer"`
	HouseBank             *string `json:"HouseBank"`
	ClearCustomerSupplier *bool   `json:"ClearCustomerSupplier"`
	ReconciliationAccount *string `json:"ReconciliationAccount"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}
