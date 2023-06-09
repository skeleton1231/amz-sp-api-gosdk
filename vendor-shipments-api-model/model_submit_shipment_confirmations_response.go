/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response schema for the SubmitShipmentConfirmations operation.
type SubmitShipmentConfirmationsResponse struct {
	Payload *TransactionReference `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
