/*
 * Selling Partner API for Retail Procurement Payments
 *
 * The Selling Partner API for Retail Procurement Payments provides programmatic access to vendors payments data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PartyIdentification struct {
	// Assigned identification for the party.
	PartyId string `json:"partyId"`
	Address *Address `json:"address,omitempty"`
	// Tax registration details of the party.
	TaxRegistrationDetails []TaxRegistrationDetails `json:"taxRegistrationDetails,omitempty"`
}
