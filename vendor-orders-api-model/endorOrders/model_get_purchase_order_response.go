/*
 * Selling Partner API for Retail Procurement Orders
 *
 * The Selling Partner API for Retail Procurement Orders provides programmatic access to vendor orders data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response schema for the getPurchaseOrder operation.
type GetPurchaseOrderResponse struct {
	Payload *Order `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
