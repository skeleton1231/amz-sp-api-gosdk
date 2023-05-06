/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Shipment measurement details.
type TransportShipmentMeasurements struct {
	// Total number of cartons present in the shipment. Provide the cartonCount only for non-palletized shipments.
	TotalCartonCount int32 `json:"totalCartonCount,omitempty"`
	// Total number of Stackable Pallets present in the shipment.
	TotalPalletStackable int32 `json:"totalPalletStackable,omitempty"`
	// Total number of Non Stackable Pallets present in the shipment.
	TotalPalletNonStackable int32 `json:"totalPalletNonStackable,omitempty"`
	ShipmentWeight *Weight `json:"shipmentWeight,omitempty"`
	ShipmentVolume *Volume `json:"shipmentVolume,omitempty"`
}