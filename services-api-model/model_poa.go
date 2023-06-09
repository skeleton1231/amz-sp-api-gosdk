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

// Proof of Appointment (POA) details.
type Poa struct {
	AppointmentTime *AppointmentTime `json:"appointmentTime,omitempty"`
	// A list of technicians.
	Technicians []Technician `json:"technicians,omitempty"`
	// The identifier of the technician who uploaded the POA.
	UploadingTechnician string `json:"uploadingTechnician,omitempty"`
	// The date and time when the POA was uploaded in ISO 8601 format.
	UploadTime time.Time `json:"uploadTime,omitempty"`
	// The type of POA uploaded.
	PoaType string `json:"poaType,omitempty"`
}
