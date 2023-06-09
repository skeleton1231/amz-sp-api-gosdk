/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// `UpdateScheduleRecord` entity contains the `AvailabilityRecord` if there is an error/warning while performing the requested operation on it.
type UpdateScheduleRecord struct {
	Availability *AvailabilityRecord `json:"availability,omitempty"`
	Warnings *[]Warning `json:"warnings,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
