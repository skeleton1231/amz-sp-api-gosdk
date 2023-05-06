/*
 * Selling Partner API for Direct Fulfillment Orders
 *
 * The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Shipment details required for the shipment.
type ShipmentDetails struct {
	// When true, this is a priority shipment.
	IsPriorityShipment bool `json:"isPriorityShipment"`
	// When true, this order is part of a scheduled delivery program.
	IsScheduledDeliveryShipment bool `json:"isScheduledDeliveryShipment,omitempty"`
	// When true, a packing slip is required to be sent to the customer.
	IsPslipRequired bool `json:"isPslipRequired"`
	// When true, the order contain a gift. Include the gift message and gift wrap information.
	IsGift bool `json:"isGift,omitempty"`
	// Ship method to be used for shipping the order. Amazon defines ship method codes indicating the shipping carrier and shipment service level. To see the full list of ship methods in use, including both the code and the friendly name, search the 'Help' section on Vendor Central for 'ship methods'.
	ShipMethod string `json:"shipMethod"`
	ShipmentDates *ShipmentDates `json:"shipmentDates"`
	// Message to customer for order status.
	MessageToCustomer string `json:"messageToCustomer"`
}
