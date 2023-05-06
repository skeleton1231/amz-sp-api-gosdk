/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items provides programmatic access to information about items in the Amazon catalog.  For more information, refer to the [Catalog Items API Use Case Guide](doc:catalog-items-api-v2022-04-01-use-case-guide).
 *
 * API version: 2022-04-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Image for an item in the Amazon catalog.
type ItemImage struct {
	// Variant of the image, such as `MAIN` or `PT01`.
	Variant string `json:"variant"`
	// Link, or URL, for the image.
	Link string `json:"link"`
	// Height of the image in pixels.
	Height int32 `json:"height"`
	// Width of the image in pixels.
	Width int32 `json:"width"`
}