/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * API version: 2020-07-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// CurrentStatus : The current delivery status of the package.
type CurrentStatus string

// List of CurrentStatus
const (
	IN_TRANSIT_CurrentStatus CurrentStatus = "IN_TRANSIT"
	DELIVERED_CurrentStatus CurrentStatus = "DELIVERED"
	RETURNING_CurrentStatus CurrentStatus = "RETURNING"
	RETURNED_CurrentStatus CurrentStatus = "RETURNED"
	UNDELIVERABLE_CurrentStatus CurrentStatus = "UNDELIVERABLE"
	DELAYED_CurrentStatus CurrentStatus = "DELAYED"
	AVAILABLE_FOR_PICKUP_CurrentStatus CurrentStatus = "AVAILABLE_FOR_PICKUP"
	CUSTOMER_ACTION_CurrentStatus CurrentStatus = "CUSTOMER_ACTION"
	UNKNOWN_CurrentStatus CurrentStatus = "UNKNOWN"
	OUT_FOR_DELIVERY_CurrentStatus CurrentStatus = "OUT_FOR_DELIVERY"
	DELIVERY_ATTEMPTED_CurrentStatus CurrentStatus = "DELIVERY_ATTEMPTED"
	PICKUP_SUCCESSFUL_CurrentStatus CurrentStatus = "PICKUP_SUCCESSFUL"
	PICKUP_CANCELLED_CurrentStatus CurrentStatus = "PICKUP_CANCELLED"
	PICKUP_ATTEMPTED_CurrentStatus CurrentStatus = "PICKUP_ATTEMPTED"
	PICKUP_SCHEDULED_CurrentStatus CurrentStatus = "PICKUP_SCHEDULED"
	RETURN_REQUEST_ACCEPTED_CurrentStatus CurrentStatus = "RETURN_REQUEST_ACCEPTED"
	REFUND_ISSUED_CurrentStatus CurrentStatus = "REFUND_ISSUED"
	RETURN_RECEIVED_IN_FC_CurrentStatus CurrentStatus = "RETURN_RECEIVED_IN_FC"
)
