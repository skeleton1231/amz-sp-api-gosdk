/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Information about withheld taxes.
type TaxCollection struct {
	// The tax collection model applied to the item.
	Model string `json:"Model,omitempty"`
	// The party responsible for withholding the taxes and remitting them to the taxing authority.
	ResponsibleParty string `json:"ResponsibleParty,omitempty"`
}