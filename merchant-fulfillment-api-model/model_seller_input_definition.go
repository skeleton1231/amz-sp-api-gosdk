/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Specifies characteristics that apply to a seller input.
type SellerInputDefinition struct {
	// When true, the additional input field is required.
	IsRequired bool `json:"IsRequired"`
	// The data type of the additional input field.
	DataType string `json:"DataType"`
	Constraints *[]Constraint `json:"Constraints"`
	// The display text for the additional input field.
	InputDisplayText string `json:"InputDisplayText"`
	InputTarget *InputTargetType `json:"InputTarget,omitempty"`
	StoredValue *AdditionalSellerInput `json:"StoredValue"`
	RestrictedSetValues *[]string `json:"RestrictedSetValues,omitempty"`
}
