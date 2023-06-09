/*
 * Selling Partner API for Finances
 *
 * The Selling Partner API for Finances helps you obtain financial information relevant to a seller's business. You can obtain financial events for a given order, financial event group, or date range without having to wait until a statement period closes. You can also obtain financial event groups for a given date range.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// A removal shipment event for a removal order.
type RemovalShipmentEvent struct {
	PostedDate *time.Time `json:"PostedDate,omitempty"`
	// The merchant removal orderId.
	MerchantOrderId string `json:"MerchantOrderId,omitempty"`
	// The identifier for the removal shipment order.
	OrderId string `json:"OrderId,omitempty"`
	// The type of removal order.  Possible values:  * WHOLESALE_LIQUIDATION
	TransactionType string `json:"TransactionType,omitempty"`
	RemovalShipmentItemList *[]RemovalShipmentItem `json:"RemovalShipmentItemList,omitempty"`
}
