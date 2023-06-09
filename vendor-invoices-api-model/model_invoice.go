/*
 * Selling Partner API for Retail Procurement Payments
 *
 * The Selling Partner API for Retail Procurement Payments provides programmatic access to vendors payments data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Invoice struct {
	// Identifies the type of invoice.
	InvoiceType string `json:"invoiceType"`
	// Unique number relating to the charges defined in this document. This will be invoice number if the document type is Invoice or CreditNote number if the document type is Credit Note. Failure to provide this reference will result in a rejection.
	Id string `json:"id"`
	// An additional unique reference number used for regulatory or other purposes.
	ReferenceNumber string `json:"referenceNumber,omitempty"`
	Date *time.Time `json:"date"`
	RemitToParty *PartyIdentification `json:"remitToParty"`
	ShipToParty *PartyIdentification `json:"shipToParty,omitempty"`
	ShipFromParty *PartyIdentification `json:"shipFromParty,omitempty"`
	BillToParty *PartyIdentification `json:"billToParty,omitempty"`
	PaymentTerms *PaymentTerms `json:"paymentTerms,omitempty"`
	InvoiceTotal *Money `json:"invoiceTotal"`
	// Total tax amount details for all line items.
	TaxDetails []TaxDetails `json:"taxDetails,omitempty"`
	// Additional details provided by the selling party, for tax related or other purposes.
	AdditionalDetails []AdditionalDetails `json:"additionalDetails,omitempty"`
	// Total charge amount details for all line items.
	ChargeDetails []ChargeDetails `json:"chargeDetails,omitempty"`
	// Total allowance amount details for all line items.
	AllowanceDetails []AllowanceDetails `json:"allowanceDetails,omitempty"`
	// The list of invoice items.
	Items []InvoiceItem `json:"items,omitempty"`
}
