/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ShippingLabel struct {
	// This field will contain the Purchase Order Number for this order.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	SellingParty *PartyIdentification `json:"sellingParty"`
	ShipFromParty *PartyIdentification `json:"shipFromParty"`
	// Format of the label.
	LabelFormat string `json:"labelFormat"`
	// Provides the details of the packages in this shipment.
	LabelData []LabelData `json:"labelData"`
}
