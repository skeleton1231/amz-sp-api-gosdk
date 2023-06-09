/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Response schema for the `getRangeSlotCapacity` operation.
type RangeSlotCapacity struct {
	// Resource Identifier.
	ResourceId string `json:"resourceId,omitempty"`
	// Array of range capacities where each entry is for a specific capacity type.
	Capacities []RangeCapacity `json:"capacities,omitempty"`
	// Next page token, if there are more pages.
	NextPageToken string `json:"nextPageToken,omitempty"`
}
