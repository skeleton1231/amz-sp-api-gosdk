/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items provides programmatic access to information about items in the Amazon catalog.  For more information, refer to the [Catalog Items API Use Case Guide](doc:catalog-items-api-v2022-04-01-use-case-guide).
 *
 * API version: 2022-04-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Description of a brand that can be used to get more fine-grained search results.
type BrandRefinement struct {
	// The estimated number of results that would still be returned if refinement key applied.
	NumberOfResults int32 `json:"numberOfResults"`
	// Brand name. For display and can be used as a search refinement.
	BrandName string `json:"brandName"`
}
