package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Customer          Customer `json:"BusinessPartnerCustomer"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Customer struct {
	BusinessPartner               int             `json:"BusinessPartner"`
	Customer                      int             `json:"Customer"`
	Currency                      *string         `json:"Currency"`
	PaymentTerms                  *string         `json:"PaymentTerms"`
	PaymentMethod                 *string         `json:"PaymentMethod"`
	Incoterms                     *string         `json:"Incoterms"`
	BPAccountAssignmentGroup      *string         `json:"BPAccountAssignmentGroup"`
	CreationDate                  *string         `json:"CreationDate"`
	QuotationIsBlockedForCustomer *bool           `json:"QuotationIsBlockedForCustomer"`
	OrderIsBlockedForCustomer     *bool           `json:"OrderIsBlockedForCustomer"`
	DeliveryIsBlockedForCustomer  *bool           `json:"DeliveryIsBlockedForCustomer"`
	BillingIsBlockedForCustomer   *bool           `json:"BillingIsBlockedForCustomer"`
	PostingIsBlockedForCustomer   *bool           `json:"PostingIsBlockedForCustomer"`
	IsMarkedForDeletion           *bool           `json:"IsMarkedForDeletion"`
	Contact                       Contact         `json:"Contact"`
	PartnerFunction               PartnerFunction `json:"PartnerFunction"`
	Tax                           Tax             `json:"Tax"`
	Accounting                    Accounting      `json:"Accounting"`
	FinInst                       FinInst         `json:"FinInst"`
}

type Contact struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	Customer            int     `json:"Customer"`
	ContactID           int     `json:"ContactID"`
	ContactPersonName   *string `json:"ContactPersonName"`
	EmailAddress        *string `json:"EmailAddress"`
	PhoneNumber         *string `json:"PhoneNumber"`
	MobilePhoneNumber   *string `json:"MobilePhoneNumber"`
	FaxNumber           *string `json:"FaxNumber"`
	ContactTag1         *string `json:"ContactTag1"`
	ContactTag2         *string `json:"ContactTag2"`
	ContactTag3         *string `json:"ContactTag3"`
	ContactTag4         *string `json:"ContactTag4"`
	DefaultContact      *bool   `json:"DefaultContact"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type PartnerFunction struct {
	BusinessPartner                int                    `json:"BusinessPartner"`
	Customer                       int                    `json:"Customer"`
	PartnerCounter                 int                    `json:"PartnerCounter"`
	PartnerFunction                *string                `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int                   `json:"PartnerFunctionBusinessPartner"`
	DefaultPartner                 *bool                  `json:"DefaultPartner"`
	CreationDate                   *string                `json:"CreationDate"`
	IsMarkedForDeletion            *bool                  `json:"IsMarkedForDeletion"`
	PartnerFunctionContact         PartnerFunctionContact `json:"PartnerFunctionContact"`
	PartnerPlant                   PartnerPlant           `json:"PartnerPlant"`
}

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

type FinInst struct {
	BusinessPartner           int     `json:"BusinessPartner"`
	Customer                  int     `json:"Customer"`
	FinInstIdentification     int     `json:"FinInstIdentification"`
	ValidityEndDate           string  `json:"ValidityEndDate"`
	ValidityStartDate         string  `json:"ValidityStartDate"`
	FinInstCountry            *string `json:"FinInstCountry"`
	FinInstNumber             *string `json:"FinInstNumber"`
	FinInstName               *string `json:"FinInstName"`
	FinInstBranchName         *string `json:"FinInstBranchName"`
	SWIFTCode                 *string `json:"SWIFTCode"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID"`
	FinInstControlKey         *string `json:"FinInstControlKey"`
	FinInstAccountName        *string `json:"FinInstAccountName"`
	FinInstAccount            *string `json:"FinInstAccount"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type Tax struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	Customer            int     `json:"Customer"`
	DepartureCountry    string  `json:"DepartureCountry"`
	TaxCategory         *string `json:"TaxCategory"`
	BPTaxClassification *string `json:"BPTaxClassification"`
}

type Accounting struct {
	BusinessPartner       int     `json:"BusinessPartner"`
	Customer              int     `json:"Customer"`
	HouseBank             *string `json:"HouseBank"`
	ClearCustomerSupplier *bool   `json:"ClearCustomerSupplier"`
	ReconciliationAccount *string `json:"ReconciliationAccount"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}
