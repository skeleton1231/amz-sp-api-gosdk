/*
 * Selling Partner API for Vendor Direct Fulfillment Sandbox Test Data
 *
 * The Selling Partner API for Vendor Direct Fulfillment Sandbox Test Data provides programmatic access to vendor direct fulfillment sandbox test data.
 *
 * API version: 2021-10-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A scenario test case response returned when the request is successful.
type Scenario struct {
	// An identifier that identifies the type of scenario that user can use for testing.
	ScenarioId string `json:"scenarioId"`
	// A list of orders that can be used by the caller to test each life cycle or scenario.
	Orders []TestOrder `json:"orders"`
}
