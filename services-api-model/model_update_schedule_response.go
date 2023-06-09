/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Response schema for the `updateSchedule` operation.
type UpdateScheduleResponse struct {
	// Contains the `UpdateScheduleRecords` for which the error/warning has occurred.
	Payload []UpdateScheduleRecord `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
