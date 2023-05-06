/*
 * Selling Partner API for A+ Content Management
 *
 * With the A+ Content API, you can build applications that help selling partners add rich marketing content to their Amazon product detail pages. A+ content helps selling partners share their brand and product story, which helps buyers make informed purchasing decisions. Selling partners assemble content by choosing from content modules and adding images and text.
 *
 * API version: 2020-11-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A reference to an image, hosted in the A+ Content media library.
type ImageComponent struct {
	// This identifier is provided by the Selling Partner API for Uploads.
	UploadDestinationId string `json:"uploadDestinationId"`
	ImageCropSpecification *ImageCropSpecification `json:"imageCropSpecification"`
	// The alternative text for the image.
	AltText string `json:"altText"`
}
