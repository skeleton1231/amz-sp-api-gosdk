/*
 * Selling Partner API for Direct Fulfillment Orders
 *
 * The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Total tax details for the line item.
type OrderItemTaxDetails struct {
	TaxLineItem *[]TaxDetails `json:"taxLineItem,omitempty"`
}
