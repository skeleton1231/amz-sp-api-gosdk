/*
 * Selling Partner API for Product Fees
 *
 * The Selling Partner API for Product Fees lets you programmatically retrieve estimated fees for a product. You can then account for those fees in your pricing.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An item identifier, marketplace, time of request, and other details that identify an estimate.
type FeesEstimateIdentifier struct {
	// A marketplace identifier.
	MarketplaceId string `json:"MarketplaceId,omitempty"`
	// The seller identifier.
	SellerId string `json:"SellerId,omitempty"`
	IdType *IdType `json:"IdType,omitempty"`
	// The item identifier.
	IdValue string `json:"IdValue,omitempty"`
	// When true, the offer is fulfilled by Amazon.
	IsAmazonFulfilled bool `json:"IsAmazonFulfilled,omitempty"`
	PriceToEstimateFees *PriceToEstimateFees `json:"PriceToEstimateFees,omitempty"`
	// A unique identifier provided by the caller to track this request.
	SellerInputIdentifier string `json:"SellerInputIdentifier,omitempty"`
	OptionalFulfillmentProgram *OptionalFulfillmentProgram `json:"OptionalFulfillmentProgram,omitempty"`
}