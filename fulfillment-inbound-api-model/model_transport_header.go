/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The shipping identifier, information about whether the shipment is by an Amazon-partnered carrier, and information about whether the shipment is Small Parcel or Less Than Truckload/Full Truckload (LTL/FTL).
type TransportHeader struct {
	// The Amazon seller identifier.
	SellerId string `json:"SellerId"`
	// A shipment identifier originally returned by the createInboundShipmentPlan operation.
	ShipmentId string `json:"ShipmentId"`
	// Indicates whether a putTransportDetails request is for a partnered carrier.  Possible values:  * true – Request is for an Amazon-partnered carrier.  * false – Request is for a non-Amazon-partnered carrier.
	IsPartnered bool `json:"IsPartnered"`
	ShipmentType *ShipmentType `json:"ShipmentType"`
}