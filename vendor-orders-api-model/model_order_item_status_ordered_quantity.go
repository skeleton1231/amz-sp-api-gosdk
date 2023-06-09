/*
 * Selling Partner API for Retail Procurement Orders
 *
 * The Selling Partner API for Retail Procurement Orders provides programmatic access to vendor orders data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Ordered quantity information.
type OrderItemStatusOrderedQuantity struct {
	OrderedQuantity *ItemQuantity `json:"orderedQuantity,omitempty"`
	// Details of item quantity ordered.
	OrderedQuantityDetails []OrderedQuantityDetails `json:"orderedQuantityDetails,omitempty"`
}
