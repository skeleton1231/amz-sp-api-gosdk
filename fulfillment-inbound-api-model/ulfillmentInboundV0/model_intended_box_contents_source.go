/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// IntendedBoxContentsSource : How the seller intends to provide box contents information for a shipment.
type IntendedBoxContentsSource string

// List of IntendedBoxContentsSource
const (
	NONE_IntendedBoxContentsSource IntendedBoxContentsSource = "NONE"
	FEED_IntendedBoxContentsSource IntendedBoxContentsSource = "FEED"
	2DBARCODE__IntendedBoxContentsSource IntendedBoxContentsSource = "2D_BARCODE"
)
