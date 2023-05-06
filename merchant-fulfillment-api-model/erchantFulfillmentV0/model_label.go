/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Data for creating a shipping label and dimensions for printing the label.
type Label struct {
	CustomTextForLabel string `json:"CustomTextForLabel,omitempty"`
	Dimensions *LabelDimensions `json:"Dimensions"`
	FileContents *FileContents `json:"FileContents"`
	LabelFormat *LabelFormat `json:"LabelFormat,omitempty"`
	StandardIdForLabel *StandardIdForLabel `json:"StandardIdForLabel,omitempty"`
}
