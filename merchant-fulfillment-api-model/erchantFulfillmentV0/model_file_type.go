/*
 * Selling Partner API for Merchant Fulfillment
 *
 * The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// FileType : The file type for a label.
type FileType string

// List of FileType
const (
	APPLICATIONPDF_FileType FileType = "application/pdf"
	APPLICATIONZPL_FileType FileType = "application/zpl"
	IMAGEPNG_FileType FileType = "image/png"
)
