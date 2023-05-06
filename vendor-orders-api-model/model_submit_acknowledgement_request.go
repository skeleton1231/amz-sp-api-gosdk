/*
 * Selling Partner API for Retail Procurement Orders
 *
 * The Selling Partner API for Retail Procurement Orders provides programmatic access to vendor orders data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The request schema for the submitAcknowledgment operation.
type SubmitAcknowledgementRequest struct {
	Acknowledgements []OrderAcknowledgement `json:"acknowledgements,omitempty"`
}
