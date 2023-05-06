/*
 * Selling Partner API for Direct Fulfillment Orders
 *
 * The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The request schema for the submitAcknowledgement operation.
type SubmitAcknowledgementRequest struct {
	// A list of one or more purchase orders.
	OrderAcknowledgements []OrderAcknowledgementItem `json:"orderAcknowledgements,omitempty"`
}
