/*
 * Selling Partner API for Feeds
 *
 * Effective **June 27, 2023**, the Selling Partner API for Feeds v2020-09-04 will no longer be available and all calls to it will fail. Integrations that rely on the Feeds API must migrate to [Feeds v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference) to avoid service disruption.
 *
 * API version: 2020-09-04
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Response schema.
type GetFeedDocumentResponse struct {
	Payload *FeedDocument `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
