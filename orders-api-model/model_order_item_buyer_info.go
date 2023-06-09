/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A single order item's buyer information.
type OrderItemBuyerInfo struct {
	// An Amazon-defined order item identifier.
	OrderItemId string `json:"OrderItemId"`
	BuyerCustomizedInfo *BuyerCustomizedInfoDetail `json:"BuyerCustomizedInfo,omitempty"`
	GiftWrapPrice *Money `json:"GiftWrapPrice,omitempty"`
	GiftWrapTax *Money `json:"GiftWrapTax,omitempty"`
	// A gift message provided by the buyer.
	GiftMessageText string `json:"GiftMessageText,omitempty"`
	// The gift wrap level specified by the buyer.
	GiftWrapLevel string `json:"GiftWrapLevel,omitempty"`
}
