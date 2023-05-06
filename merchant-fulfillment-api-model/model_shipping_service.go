/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// A shipping service offer made by a carrier.
type ShippingService struct {
	// A plain text representation of a carrier's shipping service. For example, \"UPS Ground\" or \"FedEx Standard Overnight\". 
	ShippingServiceName string `json:"ShippingServiceName"`
	// The name of the carrier.
	CarrierName string `json:"CarrierName"`
	ShippingServiceId string `json:"ShippingServiceId"`
	// An Amazon-defined shipping service offer identifier.
	ShippingServiceOfferId string `json:"ShippingServiceOfferId"`
	ShipDate *time.Time `json:"ShipDate"`
	EarliestEstimatedDeliveryDate *time.Time `json:"EarliestEstimatedDeliveryDate,omitempty"`
	LatestEstimatedDeliveryDate *time.Time `json:"LatestEstimatedDeliveryDate,omitempty"`
	Rate *CurrencyAmount `json:"Rate"`
	ShippingServiceOptions *ShippingServiceOptions `json:"ShippingServiceOptions"`
	AvailableShippingServiceOptions *AvailableShippingServiceOptions `json:"AvailableShippingServiceOptions,omitempty"`
	AvailableLabelFormats *[]LabelFormat `json:"AvailableLabelFormats,omitempty"`
	AvailableFormatOptionsForLabel *[]LabelFormatOption `json:"AvailableFormatOptionsForLabel,omitempty"`
	// When true, additional seller inputs are required.
	RequiresAdditionalSellerInputs bool `json:"RequiresAdditionalSellerInputs"`
}