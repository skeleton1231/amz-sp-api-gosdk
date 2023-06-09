/*
 * Selling Partner API for Listings Items
 *
 * The Selling Partner API for Listings Items (Listings Items API) provides programmatic access to selling partner listings on Amazon. Use this API in collaboration with the Selling Partner API for Product Type Definitions, which you use to retrieve the information about Amazon product types needed to use the Listings Items API.  For more information, see the [Listings Items API Use Case Guide](doc:listings-items-api-v2021-08-01-use-case-guide).
 *
 * API version: 2021-08-01
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
