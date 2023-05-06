/*
 * Selling Partner API for Pricing
 *
 * The Selling Partner API for Pricing helps you programmatically retrieve product pricing and offer information for Amazon Marketplace products.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ItemOffersRequestParams struct {
	MarketplaceId string `json:"MarketplaceId"`
	ItemCondition *ItemCondition `json:"ItemCondition"`
	CustomerType *CustomerType `json:"CustomerType,omitempty"`
	// The Amazon Standard Identification Number (ASIN) of the item. This is the same Asin passed as a request parameter.
	Asin string `json:"Asin,omitempty"`
}
