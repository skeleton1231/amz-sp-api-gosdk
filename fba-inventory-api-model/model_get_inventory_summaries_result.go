/*
 * Selling Partner API for FBA Inventory
 *
 * The Selling Partner API for FBA Inventory lets you programmatically retrieve information about inventory in Amazon's fulfillment network.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The payload schema for the getInventorySummaries operation.
type GetInventorySummariesResult struct {
	Granularity *Granularity `json:"granularity"`
	InventorySummaries *[]InventorySummary `json:"inventorySummaries"`
}
