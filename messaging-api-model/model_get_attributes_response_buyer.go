/*
 * Selling Partner API for Messaging
 *
 * With the Messaging API you can build applications that send messages to buyers. You can get a list of message types that are available for an order that you specify, then call an operation that sends a message to the buyer for that order. The Messaging API returns responses that are formed according to the <a href=https://tools.ietf.org/html/draft-kelly-json-hal-08>JSON Hypertext Application Language</a> (HAL) standard.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The list of attributes related to the buyer.
type GetAttributesResponseBuyer struct {
	// The buyer's language of preference, indicated with a locale-specific language tag. Examples: \"en-US\", \"zh-CN\", and \"en-GB\".
	Locale string `json:"locale,omitempty"`
}
