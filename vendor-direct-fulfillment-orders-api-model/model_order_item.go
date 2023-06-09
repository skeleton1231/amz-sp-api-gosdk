/*
 * Selling Partner API for Direct Fulfillment Orders
 *
 * The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OrderItem struct {
	// Numbering of the item on the purchase order. The first item will be 1, the second 2, and so on.
	ItemSequenceNumber string `json:"itemSequenceNumber"`
	// Buyer's standard identification number (ASIN) of an item.
	BuyerProductIdentifier string `json:"buyerProductIdentifier,omitempty"`
	// The vendor selected product identification of the item.
	VendorProductIdentifier string `json:"vendorProductIdentifier,omitempty"`
	// Title for the item.
	Title string `json:"title,omitempty"`
	OrderedQuantity *ItemQuantity `json:"orderedQuantity"`
	ScheduledDeliveryShipment *ScheduledDeliveryShipment `json:"scheduledDeliveryShipment,omitempty"`
	GiftDetails *GiftDetails `json:"giftDetails,omitempty"`
	NetPrice *Money `json:"netPrice"`
	TaxDetails *TaxItemDetails `json:"taxDetails,omitempty"`
	TotalPrice *Money `json:"totalPrice,omitempty"`
	BuyerCustomizedInfo *BuyerCustomizedInfoDetail `json:"buyerCustomizedInfo,omitempty"`
}
