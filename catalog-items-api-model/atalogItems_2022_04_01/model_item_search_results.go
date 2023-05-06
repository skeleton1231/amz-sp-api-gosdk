/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items provides programmatic access to information about items in the Amazon catalog.  For more information, refer to the [Catalog Items API Use Case Guide](doc:catalog-items-api-v2022-04-01-use-case-guide).
 *
 * API version: 2022-04-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Items in the Amazon catalog and search related metadata.
type ItemSearchResults struct {
	// For `identifiers`-based searches, the total number of Amazon catalog items found. For `keywords`-based searches, the estimated total number of Amazon catalog items matched by the search query (only results up to the page count limit will be returned per request regardless of the number found).  Note: The maximum number of items (ASINs) that can be returned and paged through is 1000.
	NumberOfResults int32 `json:"numberOfResults"`
	Pagination *Pagination `json:"pagination"`
	Refinements *Refinements `json:"refinements"`
	// A list of items from the Amazon catalog.
	Items []Item `json:"items"`
}
