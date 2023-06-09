/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Request schema for the `getFixedSlotCapacity` operation. This schema is used to define the time range, capacity types and slot duration which are being queried.
type FixedSlotCapacityQuery struct {
	// An array of capacity types which are being requested. Default value is `[SCHEDULED_CAPACITY]`.
	CapacityTypes []CapacityType `json:"capacityTypes,omitempty"`
	// Size in which slots are being requested. This value should be a multiple of 5 and fall in the range: 5 <= `slotDuration` <= 360.
	SlotDuration float64 `json:"slotDuration,omitempty"`
	// Start date time from which the capacity slots are being requested in ISO 8601 format.
	StartDateTime time.Time `json:"startDateTime"`
	// End date time up to which the capacity slots are being requested in ISO 8601 format.
	EndDateTime time.Time `json:"endDateTime"`
}
