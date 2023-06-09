/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Properties of packages
type PackageDetail struct {
	PackageReferenceId string `json:"packageReferenceId"`
	// Identifies the carrier that will deliver the package. This field is required for all marketplaces, see [reference](https://developer-docs.amazon.com/sp-api/changelog/carriercode-value-required-in-shipment-confirmations-for-br-mx-ca-sg-au-in-jp-marketplaces).
	CarrierCode string `json:"carrierCode"`
	// Carrier Name that will deliver the package. Required when carrierCode is \"Others\" 
	CarrierName string `json:"carrierName,omitempty"`
	// Ship method to be used for shipping the order.
	ShippingMethod string `json:"shippingMethod,omitempty"`
	// The tracking number used to obtain tracking and delivery information.
	TrackingNumber string `json:"trackingNumber"`
	// The shipping date for the package. Must be in ISO-8601 date/time format.
	ShipDate time.Time `json:"shipDate"`
	// The unique identifier of the supply source.
	ShipFromSupplySourceId string `json:"shipFromSupplySourceId,omitempty"`
	OrderItems *[]ConfirmShipmentOrderItem `json:"orderItems"`
}
