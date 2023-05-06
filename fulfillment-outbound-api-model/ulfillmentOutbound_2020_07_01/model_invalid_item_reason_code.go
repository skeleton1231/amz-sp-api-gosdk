/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * API version: 2020-07-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// InvalidItemReasonCode : A code for why the item is invalid for return.
type InvalidItemReasonCode string

// List of InvalidItemReasonCode
const (
	INVALID_VALUES_InvalidItemReasonCode InvalidItemReasonCode = "InvalidValues"
	DUPLICATE_REQUEST_InvalidItemReasonCode InvalidItemReasonCode = "DuplicateRequest"
	NO_COMPLETED_SHIP_ITEMS_InvalidItemReasonCode InvalidItemReasonCode = "NoCompletedShipItems"
	NO_RETURNABLE_QUANTITY_InvalidItemReasonCode InvalidItemReasonCode = "NoReturnableQuantity"
)
