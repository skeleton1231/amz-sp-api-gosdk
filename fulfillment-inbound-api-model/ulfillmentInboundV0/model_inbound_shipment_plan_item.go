/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Item information used to create an inbound shipment. Returned by the createInboundShipmentPlan operation.
type InboundShipmentPlanItem struct {
	// The seller SKU of the item.
	SellerSKU string `json:"SellerSKU"`
	// Amazon's fulfillment network SKU of the item.
	FulfillmentNetworkSKU string `json:"FulfillmentNetworkSKU"`
	Quantity int32 `json:"Quantity"`
	PrepDetailsList *[]PrepDetails `json:"PrepDetailsList,omitempty"`
}
