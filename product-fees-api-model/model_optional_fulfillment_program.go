/*
 * Selling Partner API for Product Fees
 *
 * The Selling Partner API for Product Fees lets you programmatically retrieve estimated fees for a product. You can then account for those fees in your pricing.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// OptionalFulfillmentProgram : An optional enrollment program to return the estimated fees when the offer is fulfilled by Amazon (IsAmazonFulfilled is set to true).
type OptionalFulfillmentProgram string

// List of OptionalFulfillmentProgram
const (
	CORE_OptionalFulfillmentProgram OptionalFulfillmentProgram = "FBA_CORE"
	SNL_OptionalFulfillmentProgram OptionalFulfillmentProgram = "FBA_SNL"
	EFN_OptionalFulfillmentProgram OptionalFulfillmentProgram = "FBA_EFN"
)
