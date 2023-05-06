/*
 * Selling Partner API for A+ Content Management
 *
 * With the A+ Content API, you can build applications that help selling partners add rich marketing content to their Amazon product detail pages. A+ content helps selling partners share their brand and product story, which helps buyers make informed purchasing decisions. Selling partners assemble content by choosing from content modules and adding images and text.
 *
 * API version: 2020-11-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Three standard images with text, presented across a single row.
type StandardThreeImageTextModule struct {
	Headline *TextComponent `json:"headline,omitempty"`
	Block1 *StandardImageTextBlock `json:"block1,omitempty"`
	Block2 *StandardImageTextBlock `json:"block2,omitempty"`
	Block3 *StandardImageTextBlock `json:"block3,omitempty"`
}
