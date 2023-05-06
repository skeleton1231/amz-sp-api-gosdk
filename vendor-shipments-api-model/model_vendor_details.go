/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Vendor Details as part of Label response.
type VendorDetails struct {
	SellingParty *PartyIdentification `json:"sellingParty,omitempty"`
	// Unique vendor shipment id which is not used in last 365 days
	VendorShipmentId time.Time `json:"vendorShipmentId,omitempty"`
}