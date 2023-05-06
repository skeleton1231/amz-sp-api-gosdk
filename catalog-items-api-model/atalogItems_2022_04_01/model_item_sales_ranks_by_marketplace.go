/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items provides programmatic access to information about items in the Amazon catalog.  For more information, refer to the [Catalog Items API Use Case Guide](doc:catalog-items-api-v2022-04-01-use-case-guide).
 *
 * API version: 2022-04-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Sales ranks of an Amazon catalog item for the indicated Amazon marketplace.
type ItemSalesRanksByMarketplace struct {
	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
	// Sales ranks of an Amazon catalog item for an Amazon marketplace by classification.
	ClassificationRanks []ItemClassificationSalesRank `json:"classificationRanks,omitempty"`
	// Sales ranks of an Amazon catalog item for an Amazon marketplace by website display group.
	DisplayGroupRanks []ItemDisplayGroupSalesRank `json:"displayGroupRanks,omitempty"`
}
