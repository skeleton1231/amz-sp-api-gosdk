/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// OtherDeliveryAttributes : Miscellaneous delivery attributes associated with the shipping address.
type OtherDeliveryAttributes string

// List of OtherDeliveryAttributes
const (
	HAS_ACCESS_POINT_OtherDeliveryAttributes OtherDeliveryAttributes = "HAS_ACCESS_POINT"
	PALLET_ENABLED_OtherDeliveryAttributes OtherDeliveryAttributes = "PALLET_ENABLED"
	PALLET_DISABLED_OtherDeliveryAttributes OtherDeliveryAttributes = "PALLET_DISABLED"
)
