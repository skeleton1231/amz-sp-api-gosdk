/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Error response returned when the request is unsuccessful.
type ModelError struct {
	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`
	// A message that describes the error condition.
	Message string `json:"message"`
	// Additional details that can help the caller understand or fix the issue.
	Details string `json:"details,omitempty"`
}