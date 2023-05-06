/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Item details for be provided for every item in shipment at either the item or carton or pallet level, whichever is appropriate.
type PurchaseOrderItemDetails struct {
	MaximumRetailPrice *Money `json:"maximumRetailPrice,omitempty"`
}
