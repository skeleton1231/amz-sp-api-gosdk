/*
 * Selling Partner API for Direct Fulfillment Orders
 *
 * The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Total tax details for the line item.
type TaxItemDetails struct {
	TaxLineItem *[]TaxDetails `json:"taxLineItem,omitempty"`
}
