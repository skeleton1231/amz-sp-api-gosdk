/*
 * Selling Partner API for Vendor Direct Fulfillment Sandbox Test Data
 *
 * The Selling Partner API for Vendor Direct Fulfillment Sandbox Test Data provides programmatic access to vendor direct fulfillment sandbox test data.
 *
 * API version: 2021-10-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The party identifiers required to generate the test data.
type OrderScenarioRequest struct {
	SellingParty *PartyIdentification `json:"sellingParty"`
	ShipFromParty *PartyIdentification `json:"shipFromParty"`
}
