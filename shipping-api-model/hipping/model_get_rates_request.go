/*
 * Selling Partner API for Shipping
 *
 * Provides programmatic access to Amazon Shipping APIs.   **Note:** If you are new to the Amazon Shipping API, refer to the latest version of <a href=\"https://developer-docs.amazon.com/amazon-shipping/docs/shipping-api-v2-reference\">Amazon Shipping API (v2)</a> on the <a href=\"https://developer-docs.amazon.com/amazon-shipping/\">Amazon Shipping Developer Documentation</a> site.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// The payload schema for the getRates operation.
type GetRatesRequest struct {
	ShipTo *Address `json:"shipTo"`
	ShipFrom *Address `json:"shipFrom"`
	ServiceTypes *[]ServiceType `json:"serviceTypes"`
	// The start date and time. This defaults to the current date and time.
	ShipDate time.Time `json:"shipDate,omitempty"`
	ContainerSpecifications *[]ContainerSpecification `json:"containerSpecifications"`
}
