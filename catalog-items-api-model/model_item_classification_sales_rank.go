/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items provides programmatic access to information about items in the Amazon catalog.  For more information, refer to the [Catalog Items API Use Case Guide](doc:catalog-items-api-v2022-04-01-use-case-guide).
 *
 * API version: 2022-04-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Sales rank of an Amazon catalog item by classification.
type ItemClassificationSalesRank struct {
	// Identifier of the classification associated with the sales rank.
	ClassificationId string `json:"classificationId"`
	// Title, or name, of the sales rank.
	Title string `json:"title"`
	// Corresponding Amazon retail website link, or URL, for the sales rank.
	Link string `json:"link,omitempty"`
	// Sales rank value.
	Rank int32 `json:"rank"`
}
