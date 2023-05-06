/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// ShipmentType : Specifies the carrier shipment type in a putTransportDetails request.
type ShipmentType string

// List of ShipmentType
const (
	SP_ShipmentType ShipmentType = "SP"
	LTL_ShipmentType ShipmentType = "LTL"
)