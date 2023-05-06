/*
 * Selling Partner API for Direct Fulfillment Payments
 *
 * The Selling Partner API for Direct Fulfillment Payments provides programmatic access to a direct fulfillment vendor's invoice data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Details of tax amount applied.
type TaxDetail struct {
	// Type of the tax applied.
	TaxType string `json:"taxType"`
	TaxRate string `json:"taxRate,omitempty"`
	TaxAmount *Money `json:"taxAmount"`
	TaxableAmount *Money `json:"taxableAmount,omitempty"`
}
