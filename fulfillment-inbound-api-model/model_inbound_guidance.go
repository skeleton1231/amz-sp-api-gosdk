/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// InboundGuidance : Specific inbound guidance for an item.
type InboundGuidance string

// List of InboundGuidance
const (
	INBOUND_NOT_RECOMMENDED_InboundGuidance InboundGuidance = "InboundNotRecommended"
	INBOUND_OK_InboundGuidance InboundGuidance = "InboundOK"
)
