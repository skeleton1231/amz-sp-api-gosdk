/*
 * Selling Partner API for Easy Ship
 *
 * The Selling Partner API for Easy Ship helps you build applications that help sellers manage and ship Amazon Easy Ship orders.  Your Easy Ship applications can:  * Get available time slots for packages to be scheduled for delivery.  * Schedule, reschedule, and cancel Easy Ship orders.  * Print labels, invoices, and warranties.  See the [Marketplace Support Table](doc:easyship-api-v2022-03-23-use-case-guide#marketplace-support-table) for the differences in Easy Ship operations by marketplace.
 *
 * API version: 2022-03-23
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The dimensions of the scheduled package.
type Dimensions struct {
	Length float32 `json:"length,omitempty"`
	Width float32 `json:"width,omitempty"`
	Height float32 `json:"height,omitempty"`
	Unit *UnitOfLength `json:"unit,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}