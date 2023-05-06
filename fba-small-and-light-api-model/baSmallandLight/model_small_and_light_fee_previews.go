/*
 * Selling Partner API for FBA Small And Light
 *
 * The Selling Partner API for FBA Small and Light lets you help sellers manage their listings in the Small and Light program. The program reduces the cost of fulfilling orders for small and lightweight FBA inventory. You can enroll or remove items from the program and check item eligibility and enrollment status. You can also preview the estimated program fees charged to a seller for items sold while enrolled in the program.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SmallAndLightFeePreviews struct {
	// A list of fee estimates for the requested items. The order of the fee estimates will follow the same order as the items in the request, with duplicates removed.
	Data []FeePreview `json:"data,omitempty"`
}
