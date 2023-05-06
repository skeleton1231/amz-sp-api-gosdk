/*
 * Selling Partner API for FBA Inventory
 *
 * The Selling Partner API for FBA Inventory lets you programmatically retrieve information about inventory in Amazon's fulfillment network.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The quantity of unfulfillable inventory.
type UnfulfillableQuantity struct {
	// The total number of units in Amazon's fulfillment network in unsellable condition.
	TotalUnfulfillableQuantity int32 `json:"totalUnfulfillableQuantity,omitempty"`
	// The number of units in customer damaged disposition.
	CustomerDamagedQuantity int32 `json:"customerDamagedQuantity,omitempty"`
	// The number of units in warehouse damaged disposition.
	WarehouseDamagedQuantity int32 `json:"warehouseDamagedQuantity,omitempty"`
	// The number of units in distributor damaged disposition.
	DistributorDamagedQuantity int32 `json:"distributorDamagedQuantity,omitempty"`
	// The number of units in carrier damaged disposition.
	CarrierDamagedQuantity int32 `json:"carrierDamagedQuantity,omitempty"`
	// The number of units in defective disposition.
	DefectiveQuantity int32 `json:"defectiveQuantity,omitempty"`
	// The number of units in expired disposition.
	ExpiredQuantity int32 `json:"expiredQuantity,omitempty"`
}