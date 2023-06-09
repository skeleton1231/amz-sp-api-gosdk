/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The regulated information collected during purchase and used to verify the order.
type RegulatedInformation struct {
	// A list of regulated information fields as collected from the regulatory form.
	Fields []RegulatedInformationField `json:"Fields"`
}
