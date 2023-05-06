/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The time interval for which the business is open.
type OpenInterval struct {
	StartTime *OpenTimeInterval `json:"StartTime,omitempty"`
	EndTime *OpenTimeInterval `json:"EndTime,omitempty"`
}
