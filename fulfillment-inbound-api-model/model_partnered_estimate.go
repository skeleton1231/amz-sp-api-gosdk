/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// The estimated shipping cost for a shipment using an Amazon-partnered carrier.
type PartneredEstimate struct {
	Amount *Amount `json:"Amount"`
	ConfirmDeadline *time.Time `json:"ConfirmDeadline,omitempty"`
	VoidDeadline *time.Time `json:"VoidDeadline,omitempty"`
}
