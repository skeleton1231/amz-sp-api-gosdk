/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * API version: 2020-07-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// An item that Amazon accepted for return.
type ReturnItem struct {
	// An identifier assigned by the seller to the return item.
	SellerReturnItemId string `json:"sellerReturnItemId"`
	// The identifier assigned to the item by the seller when the fulfillment order was created.
	SellerFulfillmentOrderItemId string `json:"sellerFulfillmentOrderItemId"`
	// The identifier for the shipment that is associated with the return item.
	AmazonShipmentId string `json:"amazonShipmentId"`
	// The return reason code assigned to the return item by the seller.
	SellerReturnReasonCode string `json:"sellerReturnReasonCode"`
	// An optional comment about the return item.
	ReturnComment string `json:"returnComment,omitempty"`
	// The return reason code that the Amazon fulfillment center assigned to the return item.
	AmazonReturnReasonCode string `json:"amazonReturnReasonCode,omitempty"`
	Status *FulfillmentReturnItemStatus `json:"status"`
	StatusChangedDate *time.Time `json:"statusChangedDate"`
	// Identifies the return authorization used to return this item. See ReturnAuthorization.
	ReturnAuthorizationId string `json:"returnAuthorizationId,omitempty"`
	ReturnReceivedCondition *ReturnItemDisposition `json:"returnReceivedCondition,omitempty"`
	// The identifier for the Amazon fulfillment center that processed the return item.
	FulfillmentCenterId string `json:"fulfillmentCenterId,omitempty"`
}
