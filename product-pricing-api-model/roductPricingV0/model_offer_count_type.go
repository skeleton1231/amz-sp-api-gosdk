/*
 * Selling Partner API for Pricing
 *
 * The Selling Partner API for Pricing helps you programmatically retrieve product pricing and offer information for Amazon Marketplace products.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The total number of offers for the specified condition and fulfillment channel.
type OfferCountType struct {
	// Indicates the condition of the item. For example: New, Used, Collectible, Refurbished, or Club.
	Condition string `json:"condition,omitempty"`
	FulfillmentChannel *FulfillmentChannelType `json:"fulfillmentChannel,omitempty"`
	// The number of offers in a fulfillment channel that meet a specific condition.
	OfferCount int32 `json:"OfferCount,omitempty"`
}
