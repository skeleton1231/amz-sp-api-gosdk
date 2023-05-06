/*
 * Selling Partner API for Catalog Items
 *
 * The Selling Partner API for Catalog Items helps you programmatically retrieve item details for items in the catalog.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SalesRankType struct {
	// Identifies the item category from which the sales rank is taken.
	ProductCategoryId string `json:"ProductCategoryId"`
	// The sales rank of the item within the item category.
	Rank int32 `json:"Rank"`
}
