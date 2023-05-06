/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ShipmentConfirmation struct {
	// Purchase order number corresponding to the shipment.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	ShipmentDetails *ShipmentDetails `json:"shipmentDetails"`
	SellingParty *PartyIdentification `json:"sellingParty"`
	ShipFromParty *PartyIdentification `json:"shipFromParty"`
	// Provide the details of the items in this shipment. If any of the item details field is common at a package or a pallet level, then provide them at the corresponding package.
	Items []Item `json:"items"`
	// Provide the details of the items in this shipment. If any of the item details field is common at a package or a pallet level, then provide them at the corresponding package.
	Containers []Container `json:"containers,omitempty"`
}