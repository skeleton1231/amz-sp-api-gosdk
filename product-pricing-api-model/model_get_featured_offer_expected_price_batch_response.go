/*
 * Selling Partner API for Pricing
 *
 * The Selling Partner API for Pricing helps you programmatically retrieve product pricing and offer pricing information for Amazon Marketplace products.  For more information, see the [Product Pricing v2022-05-01 Use Case Guide](doc:product-pricing-api-v2022-05-01-use-case-guide).
 *
 * API version: 2022-05-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response schema for the getFeaturedOfferExpectedPriceBatch operation.
type GetFeaturedOfferExpectedPriceBatchResponse struct {
	Responses *[]FeaturedOfferExpectedPriceResponse `json:"responses,omitempty"`
}
