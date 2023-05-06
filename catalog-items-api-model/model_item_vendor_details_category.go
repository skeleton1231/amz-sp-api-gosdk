/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items provides programmatic access to information about items in the Amazon catalog.  For more information, refer to the [Catalog Items API Use Case Guide](doc:catalog-items-api-v2022-04-01-use-case-guide).
 *
 * API version: 2022-04-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Product category or subcategory associated with an Amazon catalog item.
type ItemVendorDetailsCategory struct {
	// Display name of the product category or subcategory
	DisplayName string `json:"displayName,omitempty"`
	// Value (code) of the product category or subcategory.
	Value string `json:"value,omitempty"`
}
