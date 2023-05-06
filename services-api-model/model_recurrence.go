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

// Repeated occurrence of an event in a time range.
type Recurrence struct {
	// End time of the recurrence.
	EndTime time.Time `json:"endTime"`
	// Days of the week when recurrence is valid. If the schedule is valid every Monday, input will only contain `MONDAY` in the list.
	DaysOfWeek []DayOfWeek `json:"daysOfWeek,omitempty"`
	// Days of the month when recurrence is valid.
	DaysOfMonth []int32 `json:"daysOfMonth,omitempty"`
}