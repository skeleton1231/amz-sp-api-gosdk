/*
 * Selling Partner API for A+ Content Management
 *
 * With the A+ Content API, you can build applications that help selling partners add rich marketing content to their Amazon product detail pages. A+ content helps selling partners share their brand and product story, which helps buyers make informed purchasing decisions. Selling partners assemble content by choosing from content modules and adding images and text.
 *
 * API version: 2020-11-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SearchContentPublishRecordsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`
	PublishRecordList *[]PublishRecord `json:"publishRecordList"`
}
