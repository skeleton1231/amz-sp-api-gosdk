/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PartyIdentification struct {
	// Assigned Identification for the party.
	PartyId string `json:"partyId"`
	Address *Address `json:"address,omitempty"`
	// Tax registration details of the entity.
	TaxRegistrationDetails []TaxRegistrationDetails `json:"taxRegistrationDetails,omitempty"`
}
