/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items provides programmatic access to information about items in the Amazon catalog.  For more information, refer to the [Catalog Items API Use Case Guide](doc:catalog-items-api-v2022-04-01-use-case-guide).
 *
 * API version: 2022-04-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Individual contributor to the creation of an item, such as an author or actor.
type ItemContributor struct {
	Role *ItemContributorRole `json:"role"`
	// Name of the contributor, such as Jane Austen.
	Value string `json:"value"`
}