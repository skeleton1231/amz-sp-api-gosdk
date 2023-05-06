/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ShippingLabelRequest struct {
	// Purchase order number of the order for which to create a shipping label.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	SellingParty *PartyIdentification `json:"sellingParty"`
	ShipFromParty *PartyIdentification `json:"shipFromParty"`
	// A list of the packages in this shipment.
	Containers []Container `json:"containers,omitempty"`
}