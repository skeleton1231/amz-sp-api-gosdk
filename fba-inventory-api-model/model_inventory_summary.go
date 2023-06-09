/*
 * Selling Partner API for FBA Inventory
 *
 * The Selling Partner API for FBA Inventory lets you programmatically retrieve information about inventory in Amazon's fulfillment network.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Inventory summary for a specific item.
type InventorySummary struct {
	// The Amazon Standard Identification Number (ASIN) of an item.
	Asin string `json:"asin,omitempty"`
	// Amazon's fulfillment network SKU identifier.
	FnSku string `json:"fnSku,omitempty"`
	// The seller SKU of the item.
	SellerSku string `json:"sellerSku,omitempty"`
	// The condition of the item as described by the seller (for example, New Item).
	Condition string `json:"condition,omitempty"`
	InventoryDetails *InventoryDetails `json:"inventoryDetails,omitempty"`
	// The date and time that any quantity was last updated.
	LastUpdatedTime time.Time `json:"lastUpdatedTime,omitempty"`
	// The localized language product title of the item within the specific marketplace.
	ProductName string `json:"productName,omitempty"`
	// The total number of units in an inbound shipment or in Amazon fulfillment centers.
	TotalQuantity int32 `json:"totalQuantity,omitempty"`
}
