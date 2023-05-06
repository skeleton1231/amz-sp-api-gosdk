/*
 * Selling Partner API for Pricing
 *
 * The Selling Partner API for Pricing helps you programmatically retrieve product pricing and offer information for Amazon Marketplace products.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PriceType struct {
	LandedPrice *MoneyType `json:"LandedPrice,omitempty"`
	ListingPrice *MoneyType `json:"ListingPrice"`
	Shipping *MoneyType `json:"Shipping,omitempty"`
	Points *Points `json:"Points,omitempty"`
}
