/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TransportLabel struct {
	// Date on which label is created.
	LabelCreateDateTime string `json:"labelCreateDateTime,omitempty"`
	ShipmentInformation *ShipmentInformation `json:"shipmentInformation,omitempty"`
	// Indicates the label data,format and type associated .
	LabelData []LabelData `json:"labelData,omitempty"`
}
