/*
 * Selling Partner API for Pricing
 *
 * The Selling Partner API for Pricing helps you programmatically retrieve product pricing and offer pricing information for Amazon Marketplace products.  For more information, see the [Product Pricing v2022-05-01 Use Case Guide](doc:product-pricing-api-v2022-05-01-use-case-guide).
 *
 * API version: 2022-05-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An individual featured offer expected price request for a particular SKU.
type FeaturedOfferExpectedPriceRequest struct {
	MarketplaceId string `json:"marketplaceId"`
	Sku string `json:"sku"`
	// The URI associated with an individual request within a batch. For FeaturedOfferExpectedPrice, this should be '/products/pricing/2022-05-01/offer/featuredOfferExpectedPrice'.
	Uri string `json:"uri"`
	Method *HttpMethod `json:"method"`
	Body *map[string]interface{} `json:"body,omitempty"`
	Headers *map[string]string `json:"headers,omitempty"`
}
