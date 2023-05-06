/*
 * Selling Partner API for A+ Content Management
 *
 * With the A+ Content API, you can build applications that help selling partners add rich marketing content to their Amazon product detail pages. A+ content helps selling partners share their brand and product story, which helps buyers make informed purchasing decisions. Selling partners assemble content by choosing from content modules and adding images and text.
 *
 * API version: 2020-11-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// DecoratorType : The type of rich text decorator.
type DecoratorType string

// List of DecoratorType
const (
	LIST_ITEM_DecoratorType DecoratorType = "LIST_ITEM"
	LIST_ORDERED_DecoratorType DecoratorType = "LIST_ORDERED"
	LIST_UNORDERED_DecoratorType DecoratorType = "LIST_UNORDERED"
	STYLE_BOLD_DecoratorType DecoratorType = "STYLE_BOLD"
	STYLE_ITALIC_DecoratorType DecoratorType = "STYLE_ITALIC"
	STYLE_LINEBREAK_DecoratorType DecoratorType = "STYLE_LINEBREAK"
	STYLE_PARAGRAPH_DecoratorType DecoratorType = "STYLE_PARAGRAPH"
	STYLE_UNDERLINE_DecoratorType DecoratorType = "STYLE_UNDERLINE"
)
