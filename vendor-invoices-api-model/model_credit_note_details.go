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

// References required in order to process a credit note. This information is required only if InvoiceType is CreditNote.
type CreditNoteDetails struct {
	// Original Invoice Number when sending a credit note relating to an existing invoice. One Invoice only to be processed per Credit Note. This is mandatory for AP Credit Notes.
	ReferenceInvoiceNumber string `json:"referenceInvoiceNumber,omitempty"`
	// Debit Note Number as generated by Amazon. Recommended for Returns and COOP Credit Notes.
	DebitNoteNumber string `json:"debitNoteNumber,omitempty"`
	// Identifies the Returns Notice Number. Mandatory for all Returns Credit Notes.
	ReturnsReferenceNumber string `json:"returnsReferenceNumber,omitempty"`
	GoodsReturnDate *time.Time `json:"goodsReturnDate,omitempty"`
	// Identifies the Returned Merchandise Authorization ID, if generated.
	RmaId string `json:"rmaId,omitempty"`
	// Identifies the COOP reference used for COOP agreement. Failure to provide the COOP reference number or the Debit Note number may lead to a rejection of the Credit Note.
	CoopReferenceNumber string `json:"coopReferenceNumber,omitempty"`
	// Identifies the consignor reference number (VRET number), if generated by Amazon.
	ConsignorsReferenceNumber string `json:"consignorsReferenceNumber,omitempty"`
}