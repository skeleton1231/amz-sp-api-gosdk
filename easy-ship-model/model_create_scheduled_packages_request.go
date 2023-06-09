/*
 * Selling Partner API for Easy Ship
 *
 * The Selling Partner API for Easy Ship helps you build applications that help sellers manage and ship Amazon Easy Ship orders.  Your Easy Ship applications can:  * Get available time slots for packages to be scheduled for delivery.  * Schedule, reschedule, and cancel Easy Ship orders.  * Print labels, invoices, and warranties.  See the [Marketplace Support Table](doc:easyship-api-v2022-03-23-use-case-guide#marketplace-support-table) for the differences in Easy Ship operations by marketplace.
 *
 * API version: 2022-03-23
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The request body for the POST /easyShip/2022-03-23/packages/bulk API.
type CreateScheduledPackagesRequest struct {
	MarketplaceId string `json:"marketplaceId"`
	// An array allowing users to specify orders to be scheduled.
	OrderScheduleDetailsList []OrderScheduleDetails `json:"orderScheduleDetailsList"`
	LabelFormat *LabelFormat `json:"labelFormat"`
}
