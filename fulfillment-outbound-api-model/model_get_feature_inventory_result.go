/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * API version: 2020-07-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The payload for the getEligibileInventory operation.
type GetFeatureInventoryResult struct {
	// The requested marketplace.
	MarketplaceId string `json:"marketplaceId"`
	// The name of the feature.
	FeatureName string `json:"featureName"`
	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken string `json:"nextToken,omitempty"`
	// An array of SKUs eligible for this feature and the quantity available.
	FeatureSkus []FeatureSku `json:"featureSkus,omitempty"`
}