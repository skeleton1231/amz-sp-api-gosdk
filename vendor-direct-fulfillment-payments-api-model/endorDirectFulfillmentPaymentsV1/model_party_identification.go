/*
 * Selling Partner API for Direct Fulfillment Payments
 *
 * The Selling Partner API for Direct Fulfillment Payments provides programmatic access to a direct fulfillment vendor's invoice data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PartyIdentification struct {
	// Assigned Identification for the party.
	PartyId string `json:"partyId"`
	Address *Address `json:"address,omitempty"`
	// Tax registration details of the entity.
	TaxRegistrationDetails []TaxRegistrationDetail `json:"taxRegistrationDetails,omitempty"`
}
