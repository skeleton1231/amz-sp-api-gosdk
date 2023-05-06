/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The details of an appointment.
type Appointment struct {
	AppointmentId string `json:"appointmentId,omitempty"`
	// The status of the appointment.
	AppointmentStatus string `json:"appointmentStatus,omitempty"`
	AppointmentTime *AppointmentTime `json:"appointmentTime,omitempty"`
	// A list of technicians assigned to the service job.
	AssignedTechnicians []Technician `json:"assignedTechnicians,omitempty"`
	RescheduledAppointmentId string `json:"rescheduledAppointmentId,omitempty"`
	Poa *Poa `json:"poa,omitempty"`
}