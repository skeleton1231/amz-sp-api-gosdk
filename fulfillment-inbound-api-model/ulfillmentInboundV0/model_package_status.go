/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// PackageStatus : The shipment status of the package.
type PackageStatus string

// List of PackageStatus
const (
	SHIPPED_PackageStatus PackageStatus = "SHIPPED"
	IN_TRANSIT_PackageStatus PackageStatus = "IN_TRANSIT"
	DELIVERED_PackageStatus PackageStatus = "DELIVERED"
	CHECKED_IN_PackageStatus PackageStatus = "CHECKED_IN"
	RECEIVING_PackageStatus PackageStatus = "RECEIVING"
	CLOSED_PackageStatus PackageStatus = "CLOSED"
	DELETED_PackageStatus PackageStatus = "DELETED"
)
