/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response schema for the getOrderItemsBuyerInfo operation.
type GetOrderItemsBuyerInfoResponse struct {
	Payload *OrderItemsBuyerInfoList `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
