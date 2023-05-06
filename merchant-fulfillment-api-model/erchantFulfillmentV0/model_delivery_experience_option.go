/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// DeliveryExperienceOption : The delivery confirmation level.
type DeliveryExperienceOption string

// List of DeliveryExperienceOption
const (
	DELIVERY_CONFIRMATION_WITH_ADULT_SIGNATURE_DeliveryExperienceOption DeliveryExperienceOption = "DeliveryConfirmationWithAdultSignature"
	DELIVERY_CONFIRMATION_WITH_SIGNATURE_DeliveryExperienceOption DeliveryExperienceOption = "DeliveryConfirmationWithSignature"
	DELIVERY_CONFIRMATION_WITHOUT_SIGNATURE_DeliveryExperienceOption DeliveryExperienceOption = "DeliveryConfirmationWithoutSignature"
	NO_TRACKING_DeliveryExperienceOption DeliveryExperienceOption = "NoTracking"
	NO_PREFERENCE_DeliveryExperienceOption DeliveryExperienceOption = "NoPreference"
)
