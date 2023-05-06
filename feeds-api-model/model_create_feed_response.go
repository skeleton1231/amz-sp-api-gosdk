/*
 * Selling Partner API for Feeds
 *
 * The Selling Partner API for Feeds lets you upload data to Amazon on behalf of a selling partner.
 *
 * API version: 2021-06-30
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Response schema.
type CreateFeedResponse struct {
	// The identifier for the feed. This identifier is unique only in combination with a seller ID.
	FeedId string `json:"feedId"`
}
