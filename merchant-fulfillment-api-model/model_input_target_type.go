/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// InputTargetType : Indicates whether the additional seller input is at the item or shipment level.
type InputTargetType string

// List of InputTargetType
const (
	SHIPMENT_LEVEL_InputTargetType InputTargetType = "SHIPMENT_LEVEL"
	ITEM_LEVEL_InputTargetType InputTargetType = "ITEM_LEVEL"
)
