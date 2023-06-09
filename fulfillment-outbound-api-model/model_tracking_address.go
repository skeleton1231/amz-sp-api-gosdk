/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * API version: 2020-07-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Address information for tracking the package.
type TrackingAddress struct {
	// The city.
	City string `json:"city"`
	// The state.
	State string `json:"state"`
	// The country.
	Country string `json:"country"`
}
