/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Contains information regarding the Shipping Settings Automation program, such as whether the order's shipping settings were generated automatically, and what those settings are.
type AutomatedShippingSettings struct {
	// When true, this order has automated shipping settings generated by Amazon. This order could be identified as an SSA order.
	HasAutomatedShippingSettings bool `json:"HasAutomatedShippingSettings,omitempty"`
	// Auto-generated carrier for SSA orders.
	AutomatedCarrier string `json:"AutomatedCarrier,omitempty"`
	// Auto-generated ship method for SSA orders.
	AutomatedShipMethod string `json:"AutomatedShipMethod,omitempty"`
}
