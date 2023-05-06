/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// LabelPrepPreference : The preference for label preparation for an inbound shipment.
type LabelPrepPreference string

// List of LabelPrepPreference
const (
	SELLER_LABEL_LabelPrepPreference LabelPrepPreference = "SELLER_LABEL"
	AMAZON_LABEL_ONLY_LabelPrepPreference LabelPrepPreference = "AMAZON_LABEL_ONLY"
	AMAZON_LABEL_PREFERRED_LabelPrepPreference LabelPrepPreference = "AMAZON_LABEL_PREFERRED"
)