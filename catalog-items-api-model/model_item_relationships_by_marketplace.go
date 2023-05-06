/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items provides programmatic access to information about items in the Amazon catalog.  For more information, refer to the [Catalog Items API Use Case Guide](doc:catalog-items-api-v2022-04-01-use-case-guide).
 *
 * API version: 2022-04-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Relationship details for the Amazon catalog item for the indicated Amazon marketplace.
type ItemRelationshipsByMarketplace struct {
	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
	// Relationships for the item.
	Relationships []ItemRelationship `json:"relationships"`
}