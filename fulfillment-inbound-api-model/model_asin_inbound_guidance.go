/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Reasons why a given ASIN is not recommended for shipment to Amazon's fulfillment network.
type AsinInboundGuidance struct {
	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN string `json:"ASIN"`
	InboundGuidance *InboundGuidance `json:"InboundGuidance"`
	GuidanceReasonList *[]GuidanceReason `json:"GuidanceReasonList,omitempty"`
}
