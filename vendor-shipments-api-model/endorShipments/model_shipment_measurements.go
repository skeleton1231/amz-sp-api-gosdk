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
type ShipmentMeasurements struct {
	GrossShipmentWeight *Weight `json:"grossShipmentWeight,omitempty"`
	ShipmentVolume *Volume `json:"shipmentVolume,omitempty"`
	// Number of cartons present in the shipment. Provide the cartonCount only for non-palletized shipments.
	CartonCount int32 `json:"cartonCount,omitempty"`
	// Number of pallets present in the shipment. Provide the palletCount only for palletized shipments.
	PalletCount int32 `json:"palletCount,omitempty"`
}
