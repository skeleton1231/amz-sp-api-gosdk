/*
 * Selling Partner API for Direct Fulfillment Orders
 *
 * The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Details of an individual order being acknowledged.
type OrderAcknowledgementItem struct {
	// The purchase order number for this order. Formatting Notes: alpha-numeric code.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	// The vendor's order number for this order.
	VendorOrderNumber string `json:"vendorOrderNumber"`
	// The date and time when the order is acknowledged, in ISO-8601 date/time format. For example: 2018-07-16T23:00:00Z / 2018-07-16T23:00:00-05:00 / 2018-07-16T23:00:00-08:00.
	AcknowledgementDate time.Time `json:"acknowledgementDate"`
	AcknowledgementStatus *AcknowledgementStatus `json:"acknowledgementStatus"`
	SellingParty *PartyIdentification `json:"sellingParty"`
	ShipFromParty *PartyIdentification `json:"shipFromParty"`
	// Item details including acknowledged quantity.
	ItemAcknowledgements []OrderItemAcknowledgement `json:"itemAcknowledgements"`
}
