/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Details of the item being shipped.
type PurchaseOrderItems struct {
	// Item sequence number for the item. The first item will be 001, the second 002, and so on. This number is used as a reference to refer to this item from the carton or pallet level.
	ItemSequenceNumber string `json:"itemSequenceNumber"`
	// Amazon Standard Identification Number (ASIN) for a SKU
	BuyerProductIdentifier string `json:"buyerProductIdentifier,omitempty"`
	// The vendor selected product identification of the item. Should be the same as was sent in the purchase order.
	VendorProductIdentifier string `json:"vendorProductIdentifier,omitempty"`
	ShippedQuantity *ItemQuantity `json:"shippedQuantity"`
	MaximumRetailPrice *Money `json:"maximumRetailPrice,omitempty"`
}
