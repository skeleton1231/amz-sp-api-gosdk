/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response schema for the GetShipmentDetails operation.
type GetShipmentDetailsResponse struct {
	Payload *ShipmentDetails `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
