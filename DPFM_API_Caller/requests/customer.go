package requests

type Customer struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	Customer                      int     `json:"Customer"`
	Currency                      *string `json:"Currency"`
	PaymentTerms                  *string `json:"PaymentTerms"`
	PaymentMethod                 *string `json:"PaymentMethod"`
	Incoterms                     *string `json:"Incoterms"`
	BPAccountAssignmentGroup      *string `json:"BPAccountAssignmentGroup"`
	CreationDate                  *string `json:"CreationDate"`
	QuotationIsBlockedForCustomer *bool   `json:"QuotationIsBlockedForCustomer"`
	OrderIsBlockedForCustomer     *bool   `json:"OrderIsBlockedForCustomer"`
	DeliveryIsBlockedForCustomer  *bool   `json:"DeliveryIsBlockedForCustomer"`
	BillingIsBlockedForCustomer   *bool   `json:"BillingIsBlockedForCustomer"`
	PostingIsBlockedForCustomer   *bool   `json:"PostingIsBlockedForCustomer"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}
