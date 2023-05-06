/*
 * Selling Partner API for Shipping
 *
 * Provides programmatic access to Amazon Shipping APIs.   **Note:** If you are new to the Amazon Shipping API, refer to the latest version of <a href=\"https://developer-docs.amazon.com/amazon-shipping/docs/shipping-api-v2-reference\">Amazon Shipping API (v2)</a> on the <a href=\"https://developer-docs.amazon.com/amazon-shipping/\">Amazon Shipping Developer Documentation</a> site.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The specific rate purchased for the shipment, or null if unpurchased.
type AcceptedRate struct {
	TotalCharge *Currency `json:"totalCharge,omitempty"`
	BilledWeight *Weight `json:"billedWeight,omitempty"`
	ServiceType *ServiceType `json:"serviceType,omitempty"`
	Promise *ShippingPromiseSet `json:"promise,omitempty"`
}
