/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * API version: 2020-07-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Delivery information for a scheduled delivery. This is only available in the JP marketplace.
type ScheduledDeliveryInfo struct {
	// The time zone of the destination address for the fulfillment order preview. Must be an IANA time zone name. Example: Asia/Tokyo.
	DeliveryTimeZone string `json:"deliveryTimeZone"`
	DeliveryWindows *[]DeliveryWindow `json:"deliveryWindows"`
}
