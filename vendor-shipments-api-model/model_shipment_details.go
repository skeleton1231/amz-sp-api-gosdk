/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ShipmentDetails struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Shipments []Shipment `json:"shipments,omitempty"`
}