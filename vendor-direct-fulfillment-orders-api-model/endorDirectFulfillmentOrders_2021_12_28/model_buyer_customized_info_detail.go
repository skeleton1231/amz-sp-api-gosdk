/*
 * Selling Partner API for Direct Fulfillment Orders
 *
 * The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BuyerCustomizedInfoDetail struct {
	// A [Base 64](https://datatracker.ietf.org/doc/html/rfc4648#section-4) encoded URL using the UTF-8 character set. The URL provides the location of the zip file that specifies the types of customizations or configurations allowed by the vendor, along with types and ranges for the attributes of their products.
	CustomizedUrl string `json:"customizedUrl,omitempty"`
}
