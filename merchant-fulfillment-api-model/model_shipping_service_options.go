/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Extra services provided by a carrier.
type ShippingServiceOptions struct {
	DeliveryExperience *DeliveryExperienceType `json:"DeliveryExperience"`
	DeclaredValue *CurrencyAmount `json:"DeclaredValue,omitempty"`
	// When true, the carrier will pick up the package.  Note: Scheduled carrier pickup is available only using Dynamex (US), DPD (UK), and Royal Mail (UK).
	CarrierWillPickUp bool `json:"CarrierWillPickUp"`
	CarrierWillPickUpOption *CarrierWillPickUpOption `json:"CarrierWillPickUpOption,omitempty"`
	LabelFormat *LabelFormat `json:"LabelFormat,omitempty"`
}