/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Expiry struct {
	// Production, packaging or assembly date determined by the manufacturer. Its meaning is determined based on the trade item context.
	ManufacturerDate time.Time `json:"manufacturerDate,omitempty"`
	// The date that determines the limit of consumption or use of a product. Its meaning is determined based on the trade item context.
	ExpiryDate time.Time `json:"expiryDate,omitempty"`
	ExpiryAfterDuration *Duration `json:"expiryAfterDuration,omitempty"`
}
