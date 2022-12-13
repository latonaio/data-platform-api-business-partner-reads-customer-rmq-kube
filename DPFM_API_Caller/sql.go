package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-business-partner-reads-customer-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-business-partner-reads-customer-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var customer *dpfm_api_output_formatter.Customer
	var contact *dpfm_api_output_formatter.Contact
	var partnerFunction *dpfm_api_output_formatter.PartnerFunction
	var partnerFunctionContact *dpfm_api_output_formatter.PartnerFunctionContact
	var partnerPlant *dpfm_api_output_formatter.PartnerPlant
	var finInst *dpfm_api_output_formatter.FinInst
	var accounting *dpfm_api_output_formatter.Accounting
	var tax *dpfm_api_output_formatter.Tax
	for _, fn := range accepter {
		switch fn {
		case "Customer":
			func() {
				customer = c.Customer(mtx, input, output, errs, log)
			}()
		case "Contact":
			func() {
				contact = c.Contact(mtx, input, output, errs, log)
			}()
		case "PartnerFunction":
			func() {
				partnerFunction = c.PartnerFunction(mtx, input, output, errs, log)
			}()
		case "PartnerFunctionContact":
			func() {
				partnerFunctionContact = c.PartnerFunctionContact(mtx, input, output, errs, log)
			}()
		case "PartnerPlant":
			func() {
				partnerPlant = c.PartnerPlant(mtx, input, output, errs, log)
			}()
		case "FinInst":
			func() {
				finInst = c.FinInst(mtx, input, output, errs, log)
			}()
		case "Accounting":
			func() {
				accounting = c.Accounting(mtx, input, output, errs, log)
			}()
		case "Tax":
			func() {
				tax = c.Tax(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Customer:               customer,
		Contact:                contact,
		PartnerFunction:        partnerFunction,
		PartnerFunctionContact: partnerFunctionContact,
		PartnerPlant:           partnerPlant,
		FinInst:                finInst,
		Accounting:             accounting,
		Tax:                    tax,
	}

	return data
}

func (c *DPFMAPICaller) Customer(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Customer {
	businessPartner := input.Customer.BusinessPartner
	customer := input.Customer.Customer

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Customer, Currency, PaymentTerms, PaymentMethod, Incoterms, 
		BPAccountAssignmentGroup, CreationDate, QuotationIsBlockedForCustomer, OrderIsBlockedForCustomer, 
		DeliveryIsBlockedForCustomer, BillingIsBlockedForCustomer, PostingIsBlockedForCustomer, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_data
		WHERE (BusinessPartner, Customer) = (?, ?);`, businessPartner, customer,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToCustomer(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Contact(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Contact {
	businessPartner := input.Customer.BusinessPartner
	customer := input.Customer.Customer
	contactID := input.Customer.Contact.ContactID

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Customer, ContactID, ContactPersonName, EmailAddress, PhoneNumber, 
		MobilePhoneNumber, FaxNumber, ContactTag1, ContactTag2, ContactTag3, ContactTag4, DefaultContact, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_contact_data
		WHERE (BusinessPartner, Customer, ContactID) = (?, ?, ?);`, businessPartner, customer, contactID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToContact(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PartnerFunction(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PartnerFunction {
	businessPartner := input.Customer.BusinessPartner
	customer := input.Customer.Customer
	partnerCounter := input.Customer.PartnerFunction.PartnerCounter

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Customer, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner, 
		DefaultPartner, CreationDate, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_partner_function_data
		WHERE (BusinessPartner, Customer, PartnerCounter) = (?, ?, ?);`, businessPartner, customer, partnerCounter,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerFunction(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PartnerFunctionContact(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PartnerFunctionContact {
	businessPartner := input.Customer.BusinessPartner
	customer := input.Customer.Customer
	partnerCounter := input.Customer.PartnerFunction.PartnerCounter
	contactID := input.Customer.PartnerFunction.PartnerFunctionContact.ContactID

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Customer, PartnerCounter, ContactID, PartnerFunction, PartnerFunctionBusinessPartner, DefaultPartner, 
		ContactPersonName, EmailAddress, PhoneNumber, MobilePhoneNumber, FaxNumber, ContactTag1, ContactTag2, ContactTag3, ContactTag4, 
		DefaultContact, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_partner_function_contact_data
		WHERE (BusinessPartner, Customer, PartnerCounter, ContactID) = (?, ?, ?, ?);`, businessPartner, customer, partnerCounter, contactID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerFunctionContact(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PartnerPlant(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PartnerPlant {
	businessPartner := input.Customer.BusinessPartner
	customer := input.Customer.Customer
	partnerCounter := input.Customer.PartnerFunction.PartnerCounter
	partnerFunction := input.Customer.PartnerFunction.PartnerFunction
	partnerFunctionBusinessPartner := input.Customer.PartnerFunction.PartnerFunctionBusinessPartner
	plantCounter := input.Customer.PartnerFunction.PartnerPlant.PlantCounter

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Customer, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner, 
		PlantCounter, Plant, DefaultPlant, DefaultStockConfirmationPlant, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_partner_plant_data
		WHERE (BusinessPartner, Customer, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner, PlantCounter) = (?, ?, ?, ?, ?, ?);`, businessPartner, customer, partnerCounter, partnerFunction, partnerFunctionBusinessPartner, plantCounter,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerPlant(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) FinInst(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.FinInst {
	businessPartner := input.Customer.BusinessPartner
	customer := input.Customer.Customer
	finInstIdentification := input.Customer.FinInst.FinInstIdentification
	validityEndDate := input.Customer.FinInst.ValidityEndDate
	validityStartDate := input.Customer.FinInst.ValidityStartDate

	rows, err := c.db.Query(
		`SELECT BBusinessPartner, Customer, FinInstIdentification, ValidityEndDate, ValidityStartDate, FinInstCountry, 
		FinInstNumber, FinInstName, FinInstBranchName, SWIFTCode, InternalFinInstCustomerID, InternalFinInstAccountID, 
		FinInstControlKey, FinInstAccountName, FinInstAccount, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_fin_inst_data
		WHERE (BusinessPartner, Customer, FinInstIdentification, ValidityEndDate, ValidityStartDate) = (?, ?, ?, ?);`, businessPartner, customer, finInstIdentification, validityEndDate, validityStartDate,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToFinInst(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Tax(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Tax {
	businessPartner := input.Customer.BusinessPartner
	customer := input.Customer.Customer
	departureCountry := input.Customer.Tax.DepartureCountry

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Customer, DepartureCountry, TaxCategory, BPTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_tax_data
		WHERE (BusinessPartner, Customer, DepartureCountry) = (?, ?, ?);`, businessPartner, customer, departureCountry,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToTax(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Accounting(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Accounting {
	businessPartner := input.Customer.BusinessPartner
	customer := input.Customer.Customer

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Customer, HouseBank, ClearCustomerSupplier, ReconciliationAccount, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_accounting_data
		WHERE (BusinessPartner, Customer) = ?;`, businessPartner, customer,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToAccounting(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
