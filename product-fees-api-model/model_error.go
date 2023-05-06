/*
 * Selling Partner API for Product Fees
 *
 * The Selling Partner API for Product Fees lets you programmatically retrieve estimated fees for a product. You can then account for those fees in your pricing.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModelError struct {
	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`
	// A message that describes the error condition.
	Message string `json:"message"`
	// Additional information that can help the caller understand or fix the issue.
	Details string `json:"details,omitempty"`
}