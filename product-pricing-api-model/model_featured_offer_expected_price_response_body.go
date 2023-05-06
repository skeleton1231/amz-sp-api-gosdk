/*
 * Selling Partner API for Pricing
 *
 * The Selling Partner API for Pricing helps you programmatically retrieve product pricing and offer pricing information for Amazon Marketplace products.  For more information, see the [Product Pricing v2022-05-01 Use Case Guide](doc:product-pricing-api-v2022-05-01-use-case-guide).
 *
 * API version: 2022-05-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The featured offer expected price response data for a requested SKU.
type FeaturedOfferExpectedPriceResponseBody struct {
	OfferIdentifier *OfferIdentifier `json:"offerIdentifier"`
	FeaturedOfferExpectedPriceResults *[]FeaturedOfferExpectedPriceResult `json:"featuredOfferExpectedPriceResults,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
