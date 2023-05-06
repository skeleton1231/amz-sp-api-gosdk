/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response schema for the getLabels operation.
type GetLabelsResponse struct {
	Payload *LabelDownloadUrl `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
