/*
 * Selling Partner API for FBA Small And Light
 *
 * The Selling Partner API for FBA Small and Light lets you help sellers manage their listings in the Small and Light program. The program reduces the cost of fulfilling orders for small and lightweight FBA inventory. You can enroll or remove items from the program and check item eligibility and enrollment status. You can also preview the estimated program fees charged to a seller for items sold while enrolled in the program.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Error response returned when the request is unsuccessful.
type ModelError struct {
	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`
	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
	// Additional information that can help the caller understand or fix the issue.
	Details string `json:"details,omitempty"`
}
