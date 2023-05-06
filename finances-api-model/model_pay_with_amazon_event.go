/*
 * Selling Partner API for Finances
 *
 * The Selling Partner API for Finances helps you obtain financial information relevant to a seller's business. You can obtain financial events for a given order, financial event group, or date range without having to wait until a statement period closes. You can also obtain financial event groups for a given date range.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// An event related to the seller's Pay with Amazon account.
type PayWithAmazonEvent struct {
	// An order identifier that is specified by the seller.
	SellerOrderId string `json:"SellerOrderId,omitempty"`
	TransactionPostedDate *time.Time `json:"TransactionPostedDate,omitempty"`
	// The type of business object.
	BusinessObjectType string `json:"BusinessObjectType,omitempty"`
	// The sales channel for the transaction.
	SalesChannel string `json:"SalesChannel,omitempty"`
	Charge *ChargeComponent `json:"Charge,omitempty"`
	FeeList *[]FeeComponent `json:"FeeList,omitempty"`
	// The type of payment.  Possible values:  * Sales
	PaymentAmountType string `json:"PaymentAmountType,omitempty"`
	// A short description of this payment event.
	AmountDescription string `json:"AmountDescription,omitempty"`
	// The fulfillment channel.  Possible values:  * AFN - Amazon Fulfillment Network (Fulfillment by Amazon)  * MFN - Merchant Fulfillment Network (self-fulfilled)
	FulfillmentChannel string `json:"FulfillmentChannel,omitempty"`
	// The store name where the event occurred.
	StoreName string `json:"StoreName,omitempty"`
}