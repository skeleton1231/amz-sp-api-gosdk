/*
 * Selling Partner API for Direct Fulfillment Payments
 *
 * The Selling Partner API for Direct Fulfillment Payments provides programmatic access to a direct fulfillment vendor's invoice data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Tax registration details of the entity.
type TaxRegistrationDetail struct {
	// Tax registration type for the entity.
	TaxRegistrationType string `json:"taxRegistrationType,omitempty"`
	// Tax registration number for the party. For example, VAT ID.
	TaxRegistrationNumber string `json:"taxRegistrationNumber"`
	TaxRegistrationAddress *Address `json:"taxRegistrationAddress,omitempty"`
	// Tax registration message that can be used for additional tax related details.
	TaxRegistrationMessage string `json:"taxRegistrationMessage,omitempty"`
}