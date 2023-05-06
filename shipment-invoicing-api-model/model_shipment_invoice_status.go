/*
 * Selling Partner API for Shipment Invoicing
 *
 * The Selling Partner API for Shipment Invoicing helps you programmatically retrieve shipment invoice information in the Brazil marketplace for a selling partner’s Fulfillment by Amazon (FBA) orders.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// ShipmentInvoiceStatus : The shipment invoice status.
type ShipmentInvoiceStatus string

// List of ShipmentInvoiceStatus
const (
	PROCESSING_ShipmentInvoiceStatus ShipmentInvoiceStatus = "Processing"
	ACCEPTED_ShipmentInvoiceStatus ShipmentInvoiceStatus = "Accepted"
	ERRORED_ShipmentInvoiceStatus ShipmentInvoiceStatus = "Errored"
	NOT_FOUND_ShipmentInvoiceStatus ShipmentInvoiceStatus = "NotFound"
)
