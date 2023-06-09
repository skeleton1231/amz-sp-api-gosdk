/*
 * Selling Partner API for Sales
 *
 * The Selling Partner API for Sales provides APIs related to sales performance.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Error response returned when the request is unsuccessful.
type ModelError struct {
	// An error code that identifies the type of error that occured.
	Code string `json:"code"`
	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
	// Additional details that can help the caller understand or fix the issue.
	Details string `json:"details,omitempty"`
}
