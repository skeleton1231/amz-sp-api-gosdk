/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items provides programmatic access to information about items in the Amazon catalog.  For more information, refer to the [Catalog Items API Use Case Guide](doc:catalog-items-api-v2022-04-01-use-case-guide).
 *
 * API version: 2022-04-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Identifier associated with the item in the Amazon catalog, such as a UPC or EAN identifier.
type ItemIdentifier struct {
	// Type of identifier, such as UPC, EAN, or ISBN.
	IdentifierType string `json:"identifierType"`
	// Identifier.
	Identifier string `json:"identifier"`
}