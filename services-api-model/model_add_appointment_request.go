/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Input for add appointment operation.
type AddAppointmentRequest struct {
	AppointmentTime *AppointmentTimeInput `json:"appointmentTime"`
}
