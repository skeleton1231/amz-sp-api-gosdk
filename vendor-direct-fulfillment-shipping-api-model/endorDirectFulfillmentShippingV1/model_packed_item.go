/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PackedItem struct {
	// Item Sequence Number for the item. This must be the same value as sent in the order for a given item.
	ItemSequenceNumber int32 `json:"itemSequenceNumber"`
	// Buyer's Standard Identification Number (ASIN) of an item. Either buyerProductIdentifier or vendorProductIdentifier is required.
	BuyerProductIdentifier string `json:"buyerProductIdentifier,omitempty"`
	// The vendor selected product identification of the item. Should be the same as was sent in the Purchase Order, like SKU Number.
	VendorProductIdentifier string `json:"vendorProductIdentifier,omitempty"`
	PackedQuantity *ItemQuantity `json:"packedQuantity"`
}
