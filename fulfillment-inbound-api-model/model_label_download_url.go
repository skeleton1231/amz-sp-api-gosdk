/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LabelDownloadUrl struct {
	// URL to download the label for the package. Note: The URL will only be valid for 15 seconds
	DownloadURL string `json:"DownloadURL,omitempty"`
}
