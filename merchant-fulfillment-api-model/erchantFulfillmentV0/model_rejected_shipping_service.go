/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Information about a rejected shipping service
type RejectedShippingService struct {
	// The rejected shipping carrier name. e.g. USPS
	CarrierName string `json:"CarrierName"`
	// The rejected shipping service localized name. e.g. FedEx Standard Overnight
	ShippingServiceName string `json:"ShippingServiceName"`
	ShippingServiceId string `json:"ShippingServiceId"`
	// A reason code meant to be consumed programatically. e.g. CARRIER_CANNOT_SHIP_TO_POBOX
	RejectionReasonCode string `json:"RejectionReasonCode"`
	// A localized human readable description of the rejected reason.
	RejectionReasonMessage string `json:"RejectionReasonMessage,omitempty"`
}
