/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A carrier who is temporarily unavailable, most likely due to a service outage experienced by the carrier.
type TemporarilyUnavailableCarrier struct {
	// The name of the carrier.
	CarrierName string `json:"CarrierName"`
}
