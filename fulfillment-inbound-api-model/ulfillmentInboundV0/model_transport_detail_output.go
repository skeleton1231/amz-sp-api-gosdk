/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Inbound shipment information, including carrier details and shipment status.
type TransportDetailOutput struct {
	PartneredSmallParcelData *PartneredSmallParcelDataOutput `json:"PartneredSmallParcelData,omitempty"`
	NonPartneredSmallParcelData *NonPartneredSmallParcelDataOutput `json:"NonPartneredSmallParcelData,omitempty"`
	PartneredLtlData *PartneredLtlDataOutput `json:"PartneredLtlData,omitempty"`
	NonPartneredLtlData *NonPartneredLtlDataOutput `json:"NonPartneredLtlData,omitempty"`
}
