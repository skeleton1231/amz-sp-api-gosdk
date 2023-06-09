/*
 * Selling Partner API for FBA Inbound Eligibilty
 *
 * With the FBA Inbound Eligibility API, you can build applications that let sellers get eligibility previews for items before shipping them to Amazon's fulfillment centers. With this API you can find out if an item is eligible for inbound shipment to Amazon's fulfillment centers in a specific marketplace. You can also find out if an item is eligible for using the manufacturer barcode for FBA inventory tracking. Sellers can use this information to inform their decisions about which items to ship Amazon's fulfillment centers.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response object which contains the ASIN, marketplaceId if required, eligibility program, the eligibility status (boolean), and a list of ineligibility reason codes.
type ItemEligibilityPreview struct {
	// The ASIN for which eligibility was determined.
	Asin string `json:"asin"`
	// The marketplace for which eligibility was determined.
	MarketplaceId string `json:"marketplaceId,omitempty"`
	// The program for which eligibility was determined.
	Program string `json:"program"`
	// Indicates if the item is eligible for the program.
	IsEligibleForProgram bool `json:"isEligibleForProgram"`
	// Potential Ineligibility Reason Codes.
	IneligibilityReasonList []string `json:"ineligibilityReasonList,omitempty"`
}
