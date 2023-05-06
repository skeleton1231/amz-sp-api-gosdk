/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * API version: 2020-07-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type PackageTrackingDetails struct {
	// The package identifier.
	PackageNumber int32 `json:"packageNumber"`
	// The tracking number for the package.
	TrackingNumber string `json:"trackingNumber,omitempty"`
	// Link on swiship.com that allows customers to track the package.
	CustomerTrackingLink string `json:"customerTrackingLink,omitempty"`
	// The name of the carrier.
	CarrierCode string `json:"carrierCode,omitempty"`
	// The phone number of the carrier.
	CarrierPhoneNumber string `json:"carrierPhoneNumber,omitempty"`
	// The URL of the carrier's website.
	CarrierURL string `json:"carrierURL,omitempty"`
	ShipDate *time.Time `json:"shipDate,omitempty"`
	EstimatedArrivalDate *time.Time `json:"estimatedArrivalDate,omitempty"`
	ShipToAddress *TrackingAddress `json:"shipToAddress,omitempty"`
	CurrentStatus *CurrentStatus `json:"currentStatus,omitempty"`
	// Description corresponding to the CurrentStatus value.
	CurrentStatusDescription string `json:"currentStatusDescription,omitempty"`
	// The name of the person who signed for the package.
	SignedForBy string `json:"signedForBy,omitempty"`
	AdditionalLocationInfo *AdditionalLocationInfo `json:"additionalLocationInfo,omitempty"`
	TrackingEvents *[]TrackingEvent `json:"trackingEvents,omitempty"`
}
