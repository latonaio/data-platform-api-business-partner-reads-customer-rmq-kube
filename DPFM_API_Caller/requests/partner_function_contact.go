package requests

type PartnerFunctionContact struct {
	BusinessPartner                int     `json:"BusinessPartner"`
	Customer                       int     `json:"Customer"`
	PartnerCounter                 int     `json:"PartnerCounter"`
	ContactID                      int     `json:"ContactID"`
	PartnerFunction                *string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
	DefaultPartner                 *bool   `json:"DefaultPartner"`
	ContactPersonName              *string `json:"ContactPersonName"`
	EmailAddress                   *string `json:"EmailAddress"`
	PhoneNumber                    *string `json:"PhoneNumber"`
	MobilePhoneNumber              *string `json:"MobilePhoneNumber"`
	FaxNumber                      *string `json:"FaxNumber"`
	ContactTag1                    *string `json:"ContactTag1"`
	ContactTag2                    *string `json:"ContactTag2"`
	ContactTag3                    *string `json:"ContactTag3"`
	ContactTag4                    *string `json:"ContactTag4"`
	DefaultContact                 *bool   `json:"DefaultContact"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}
