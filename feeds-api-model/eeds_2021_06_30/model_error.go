/*
 * Selling Partner API for Feeds
 *
 * The Selling Partner API for Feeds lets you upload data to Amazon on behalf of a selling partner.
 *
 * API version: 2021-06-30
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An error response returned when the request is unsuccessful.
type ModelError struct {
	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`
	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
	// Additional details that can help the caller understand or fix the issue.
	Details string `json:"details,omitempty"`
}
