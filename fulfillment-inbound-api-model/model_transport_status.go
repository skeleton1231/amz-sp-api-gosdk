/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// TransportStatus : Indicates the status of the Amazon-partnered carrier shipment.
type TransportStatus string

// List of TransportStatus
const (
	WORKING_TransportStatus TransportStatus = "WORKING"
	ESTIMATING_TransportStatus TransportStatus = "ESTIMATING"
	ESTIMATED_TransportStatus TransportStatus = "ESTIMATED"
	ERROR_ON_ESTIMATING_TransportStatus TransportStatus = "ERROR_ON_ESTIMATING"
	CONFIRMING_TransportStatus TransportStatus = "CONFIRMING"
	CONFIRMED_TransportStatus TransportStatus = "CONFIRMED"
	ERROR_ON_CONFIRMING_TransportStatus TransportStatus = "ERROR_ON_CONFIRMING"
	VOIDING_TransportStatus TransportStatus = "VOIDING"
	VOIDED_TransportStatus TransportStatus = "VOIDED"
	ERROR_IN_VOIDING_TransportStatus TransportStatus = "ERROR_IN_VOIDING"
	ERROR__TransportStatus TransportStatus = "ERROR"
)