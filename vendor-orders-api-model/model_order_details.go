/*
 * Selling Partner API for Retail Procurement Orders
 *
 * The Selling Partner API for Retail Procurement Orders provides programmatic access to vendor orders data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Details of an order.
type OrderDetails struct {
	// The date the purchase order was placed. Must be in ISO-8601 date/time format.
	PurchaseOrderDate time.Time `json:"purchaseOrderDate"`
	// The date when purchase order was last changed by Amazon after the order was placed. This date will be greater than 'purchaseOrderDate'. This means the PO data was changed on that date and vendors are required to fulfill the  updated PO. The PO changes can be related to Item Quantity, Ship to Location, Ship Window etc. This field will not be present in orders that have not changed after creation. Must be in ISO-8601 date/time format.
	PurchaseOrderChangedDate time.Time `json:"purchaseOrderChangedDate,omitempty"`
	// The date when current purchase order state was changed. Current purchase order state is available in the field 'purchaseOrderState'. Must be in ISO-8601 date/time format.
	PurchaseOrderStateChangedDate time.Time `json:"purchaseOrderStateChangedDate"`
	// Type of purchase order.
	PurchaseOrderType string `json:"purchaseOrderType,omitempty"`
	ImportDetails *ImportDetails `json:"importDetails,omitempty"`
	// If requested by the recipient, this field will contain a promotional/deal number. The discount code line is optional. It is used to obtain a price discount on items on the order.
	DealCode string `json:"dealCode,omitempty"`
	// Payment method used.
	PaymentMethod string `json:"paymentMethod,omitempty"`
	BuyingParty *PartyIdentification `json:"buyingParty,omitempty"`
	SellingParty *PartyIdentification `json:"sellingParty,omitempty"`
	ShipToParty *PartyIdentification `json:"shipToParty,omitempty"`
	BillToParty *PartyIdentification `json:"billToParty,omitempty"`
	ShipWindow string `json:"shipWindow,omitempty"`
	DeliveryWindow string `json:"deliveryWindow,omitempty"`
	// A list of items in this purchase order.
	Items []OrderItem `json:"items"`
}
