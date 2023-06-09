/*
 * Selling Partner API for FBA Inventory
 *
 * The Selling Partner API for FBA Inventory lets you programmatically retrieve information about inventory in Amazon's fulfillment network.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The number of misplaced or warehouse damaged units that are actively being confirmed at our fulfillment centers.
type ResearchingQuantity struct {
	// The total number of units currently being researched in Amazon's fulfillment network.
	TotalResearchingQuantity int32 `json:"totalResearchingQuantity,omitempty"`
	// A list of quantity details for items currently being researched.
	ResearchingQuantityBreakdown []ResearchingQuantityEntry `json:"researchingQuantityBreakdown,omitempty"`
}
