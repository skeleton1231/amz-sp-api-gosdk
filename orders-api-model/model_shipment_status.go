/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// ShipmentStatus : The shipment status to apply.
type ShipmentStatus string

// List of ShipmentStatus
const (
	READY_FOR_PICKUP_ShipmentStatus ShipmentStatus = "ReadyForPickup"
	PICKED_UP_ShipmentStatus ShipmentStatus = "PickedUp"
	REFUSED_PICKUP_ShipmentStatus ShipmentStatus = "RefusedPickup"
)
