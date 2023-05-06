/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Tax information about the buyer.
type BuyerTaxInfo struct {
	// The legal name of the company.
	CompanyLegalName string `json:"CompanyLegalName,omitempty"`
	// The country or region imposing the tax.
	TaxingRegion string `json:"TaxingRegion,omitempty"`
	// A list of tax classifications that apply to the order.
	TaxClassifications []TaxClassification `json:"TaxClassifications,omitempty"`
}