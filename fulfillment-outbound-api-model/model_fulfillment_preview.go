/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * API version: 2020-07-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Information about a fulfillment order preview, including delivery and fee information based on shipping method.
type FulfillmentPreview struct {
	ShippingSpeedCategory *ShippingSpeedCategory `json:"shippingSpeedCategory"`
	ScheduledDeliveryInfo *ScheduledDeliveryInfo `json:"scheduledDeliveryInfo,omitempty"`
	// When true, this fulfillment order preview is fulfillable.
	IsFulfillable bool `json:"isFulfillable"`
	// When true, this fulfillment order preview is for COD (Cash On Delivery).
	IsCODCapable bool `json:"isCODCapable"`
	EstimatedShippingWeight *Weight `json:"estimatedShippingWeight,omitempty"`
	EstimatedFees *[]Fee `json:"estimatedFees,omitempty"`
	FulfillmentPreviewShipments *[]FulfillmentPreviewShipment `json:"fulfillmentPreviewShipments,omitempty"`
	UnfulfillablePreviewItems *[]UnfulfillablePreviewItem `json:"unfulfillablePreviewItems,omitempty"`
	OrderUnfulfillableReasons *[]string `json:"orderUnfulfillableReasons,omitempty"`
	// The marketplace the fulfillment order is placed against.
	MarketplaceId string `json:"marketplaceId"`
	// A list of features and their fulfillment policies to apply to the order.
	FeatureConstraints []FeatureSettings `json:"featureConstraints,omitempty"`
}
