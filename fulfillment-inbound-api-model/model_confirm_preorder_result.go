/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ConfirmPreorderResult struct {
	ConfirmedNeedByDate string `json:"ConfirmedNeedByDate,omitempty"`
	ConfirmedFulfillableDate string `json:"ConfirmedFulfillableDate,omitempty"`
}