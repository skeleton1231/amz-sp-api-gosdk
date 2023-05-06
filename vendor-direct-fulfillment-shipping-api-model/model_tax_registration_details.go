/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Tax registration details of the entity.
type TaxRegistrationDetails struct {
	// Tax registration type for the entity.
	TaxRegistrationType string `json:"taxRegistrationType,omitempty"`
	// Tax registration number for the party. For example, VAT ID.
	TaxRegistrationNumber string `json:"taxRegistrationNumber"`
	TaxRegistrationAddress *Address `json:"taxRegistrationAddress,omitempty"`
	// Tax registration message that can be used for additional tax related details.
	TaxRegistrationMessages string `json:"taxRegistrationMessages,omitempty"`
}