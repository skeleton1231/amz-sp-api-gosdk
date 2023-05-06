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

type TransportationDetails struct {
	// The type of shipment.
	ShipMode string `json:"shipMode,omitempty"`
	// The mode of transportation for this shipment.
	TransportationMode string `json:"transportationMode,omitempty"`
	// Date when shipment is performed by the Vendor to Buyer
	ShippedDate time.Time `json:"shippedDate,omitempty"`
	// Estimated Date on which shipment will be delivered from Vendor to Buyer
	EstimatedDeliveryDate time.Time `json:"estimatedDeliveryDate,omitempty"`
	// Date on which shipment will be delivered from Vendor to Buyer
	ShipmentDeliveryDate time.Time `json:"shipmentDeliveryDate,omitempty"`
	CarrierDetails *CarrierDetails `json:"carrierDetails,omitempty"`
	// Bill Of Lading (BOL) number is the unique number assigned by the vendor. The BOL present in the Shipment Confirmation message ideally matches the paper BOL provided with the shipment, but that is no must. Instead of BOL, an alternative reference number (like Delivery Note Number) for the shipment can also be sent in this field.
	BillOfLadingNumber string `json:"billOfLadingNumber,omitempty"`
}
