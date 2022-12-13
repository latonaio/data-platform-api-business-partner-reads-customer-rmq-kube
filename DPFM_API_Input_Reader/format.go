package dpfm_api_input_reader

import (
	"data-platform-api-business-partner-reads-customer-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToCustomer() *requests.Customer {
	data := sdc.Customer
	return &requests.Customer{
		BusinessPartner:               data.BusinessPartner,
		Customer:                      data.Customer,
		Currency:                      data.Currency,
		PaymentTerms:                  data.PaymentTerms,
		PaymentMethod:                 data.PaymentMethod,
		Incoterms:                     data.Incoterms,
		BPAccountAssignmentGroup:      data.BPAccountAssignmentGroup,
		CreationDate:                  data.CreationDate,
		QuotationIsBlockedForCustomer: data.QuotationIsBlockedForCustomer,
		OrderIsBlockedForCustomer:     data.OrderIsBlockedForCustomer,
		DeliveryIsBlockedForCustomer:  data.DeliveryIsBlockedForCustomer,
		BillingIsBlockedForCustomer:   data.BillingIsBlockedForCustomer,
		PostingIsBlockedForCustomer:   data.PostingIsBlockedForCustomer,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToContact() *requests.Contact {
	dataCustomer := sdc.Customer
	data := sdc.Customer.Contact
	return &requests.Contact{
		BusinessPartner:     dataCustomer.BusinessPartner,
		Customer:            dataCustomer.Customer,
		ContactID:           data.ContactID,
		ContactPersonName:   data.ContactPersonName,
		EmailAddress:        data.EmailAddress,
		PhoneNumber:         data.PhoneNumber,
		MobilePhoneNumber:   data.MobilePhoneNumber,
		FaxNumber:           data.FaxNumber,
		ContactTag1:         data.ContactTag1,
		ContactTag2:         data.ContactTag2,
		ContactTag3:         data.ContactTag3,
		ContactTag4:         data.ContactTag4,
		DefaultContact:      data.DefaultContact,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToPartnerFunction() *requests.PartnerFunction {
	dataCustomer := sdc.Customer
	data := sdc.Customer.PartnerFunction
	return &requests.PartnerFunction{
		BusinessPartner:                dataCustomer.BusinessPartner,
		Customer:                       dataCustomer.Customer,
		PartnerCounter:                 data.PartnerCounter,
		PartnerFunction:                data.PartnerFunction,
		PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		DefaultPartner:                 data.DefaultPartner,
		CreationDate:                   data.CreationDate,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToPartnerFunctionContact() *requests.PartnerFunctionContact {
	dataCustomer := sdc.Customer
	dataPartnerFunction := sdc.Customer.PartnerFunction
	data := sdc.Customer.PartnerFunction.PartnerFunctionContact
	return &requests.PartnerFunctionContact{
		BusinessPartner:                dataCustomer.BusinessPartner,
		Customer:                       dataCustomer.Customer,
		PartnerCounter:                 dataPartnerFunction.PartnerCounter,
		ContactID:                      data.ContactID,
		PartnerFunction:                data.PartnerFunction,
		PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		DefaultPartner:                 data.DefaultPartner,
		ContactPersonName:              data.ContactPersonName,
		EmailAddress:                   data.EmailAddress,
		PhoneNumber:                    data.PhoneNumber,
		MobilePhoneNumber:              data.MobilePhoneNumber,
		FaxNumber:                      data.FaxNumber,
		ContactTag1:                    data.ContactTag1,
		ContactTag2:                    data.ContactTag2,
		ContactTag3:                    data.ContactTag3,
		ContactTag4:                    data.ContactTag4,
		DefaultContact:                 data.DefaultContact,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToPartnerPlant() *requests.PartnerPlant {
	dataCustomer := sdc.Customer
	dataPartnerFunction := sdc.Customer.PartnerFunction
	data := sdc.Customer.PartnerFunction.PartnerPlant
	return &requests.PartnerPlant{
		BusinessPartner:                dataCustomer.BusinessPartner,
		Customer:                       dataCustomer.Customer,
		PartnerCounter:                 dataPartnerFunction.PartnerCounter,
		PartnerFunction:                data.PartnerFunction,
		PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		PlantCounter:                   data.PlantCounter,
		Plant:                          data.Plant,
		DefaultPlant:                   data.DefaultPlant,
		DefaultStockConfirmationPlant:  data.DefaultStockConfirmationPlant,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToFinInst() *requests.FinInst {
	dataCustomer := sdc.Customer
	data := sdc.Customer.FinInst
	return &requests.FinInst{
		BusinessPartner:           dataCustomer.BusinessPartner,
		Customer:                  dataCustomer.Customer,
		FinInstIdentification:     data.FinInstIdentification,
		ValidityEndDate:           data.ValidityEndDate,
		ValidityStartDate:         data.ValidityStartDate,
		FinInstCountry:            data.FinInstCountry,
		FinInstNumber:             data.FinInstNumber,
		FinInstName:               data.FinInstName,
		FinInstBranchName:         data.FinInstBranchName,
		SWIFTCode:                 data.SWIFTCode,
		InternalFinInstCustomerID: data.InternalFinInstCustomerID,
		InternalFinInstAccountID:  data.InternalFinInstAccountID,
		FinInstControlKey:         data.FinInstControlKey,
		FinInstAccountName:        data.FinInstAccountName,
		FinInstAccount:            data.FinInstAccount,
		IsMarkedForDeletion:       data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToTax() *requests.Tax {
	dataCustomer := sdc.Customer
	data := sdc.Customer.Tax
	return &requests.Tax{
		BusinessPartner:     dataCustomer.BusinessPartner,
		Customer:            dataCustomer.Customer,
		DepartureCountry:    data.DepartureCountry,
		TaxCategory:         data.TaxCategory,
		BPTaxClassification: data.BPTaxClassification,
	}
}

func (sdc *SDC) ConvertToAccounting() *requests.Accounting {
	dataCustomer := sdc.Customer
	data := sdc.Customer.Accounting
	return &requests.Accounting{
		BusinessPartner:       dataCustomer.BusinessPartner,
		Customer:              dataCustomer.Customer,
		HouseBank:             data.HouseBank,
		ClearCustomerSupplier: data.ClearCustomerSupplier,
		ReconciliationAccount: data.ReconciliationAccount,
		IsMarkedForDeletion:   data.IsMarkedForDeletion,
	}
}
