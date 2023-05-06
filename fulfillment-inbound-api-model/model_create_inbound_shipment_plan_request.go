/*
 * Selling Partner API for Fulfillment Inbound
 *
 * The Selling Partner API for Fulfillment Inbound lets you create applications that create and update inbound shipments of inventory to Amazon's fulfillment network.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The request schema for the createInboundShipmentPlan operation.
type CreateInboundShipmentPlanRequest struct {
	ShipFromAddress *Address `json:"ShipFromAddress"`
	LabelPrepPreference *LabelPrepPreference `json:"LabelPrepPreference"`
	// The two-character country code for the country where the inbound shipment is to be sent.  Note: Not required. Specifying both ShipToCountryCode and ShipToCountrySubdivisionCode returns an error.   Values:   ShipToCountryCode values for North America:  * CA – Canada  * MX - Mexico  * US - United States  ShipToCountryCode values for MCI sellers in Europe:  * DE – Germany  * ES – Spain  * FR – France  * GB – United Kingdom  * IT – Italy  Default: The country code for the seller's home marketplace.
	ShipToCountryCode string `json:"ShipToCountryCode,omitempty"`
	// The two-character country code, followed by a dash and then up to three characters that represent the subdivision of the country where the inbound shipment is to be sent. For example, \"IN-MH\". In full ISO 3166-2 format.  Note: Not required. Specifying both ShipToCountryCode and ShipToCountrySubdivisionCode returns an error.
	ShipToCountrySubdivisionCode string `json:"ShipToCountrySubdivisionCode,omitempty"`
	InboundShipmentPlanRequestItems *[]InboundShipmentPlanRequestItem `json:"InboundShipmentPlanRequestItems"`
}
