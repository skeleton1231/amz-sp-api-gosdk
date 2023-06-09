/*
 * Selling Partner API for Shipping
 *
 * Provides programmatic access to Amazon Shipping APIs.   **Note:** If you are new to the Amazon Shipping API, refer to the latest version of <a href=\"https://developer-docs.amazon.com/amazon-shipping/docs/shipping-api-v2-reference\">Amazon Shipping API (v2)</a> on the <a href=\"https://developer-docs.amazon.com/amazon-shipping/\">Amazon Shipping Developer Documentation</a> site.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A set of measurements for a three-dimensional object.
type Dimensions struct {
	// The length of the container.
	Length float64 `json:"length"`
	// The width of the container.
	Width float64 `json:"width"`
	// The height of the container.
	Height float64 `json:"height"`
	// The unit of these measurements.
	Unit string `json:"unit"`
}
