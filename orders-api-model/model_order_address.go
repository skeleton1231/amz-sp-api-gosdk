/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The shipping address for the order.
type OrderAddress struct {
	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`
	// Company name of the destination address.
	BuyerCompanyName string `json:"BuyerCompanyName,omitempty"`
	ShippingAddress *Address `json:"ShippingAddress,omitempty"`
	DeliveryPreferences *DeliveryPreferences `json:"DeliveryPreferences,omitempty"`
}