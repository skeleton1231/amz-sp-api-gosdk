/*
 * Selling Partner API for Orders
 *
 * The Selling Partner API for Orders helps you programmatically retrieve order information. These APIs let you develop fast, flexible, custom applications in areas like order synchronization, order research, and demand-based decision support tools.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The time when the business opens or closes.
type OpenTimeInterval struct {
	// The hour when the business opens or closes.
	Hour int32 `json:"Hour,omitempty"`
	// The minute when the business opens or closes.
	Minute int32 `json:"Minute,omitempty"`
}
