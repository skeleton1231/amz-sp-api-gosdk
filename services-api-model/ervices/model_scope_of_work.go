/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The scope of work for the order.
type ScopeOfWork struct {
	// The Amazon Standard Identification Number (ASIN) of the service job.
	Asin string `json:"asin,omitempty"`
	// The title of the service job.
	Title string `json:"title,omitempty"`
	// The number of service jobs.
	Quantity int32 `json:"quantity,omitempty"`
	// A list of skills required to perform the job.
	RequiredSkills []string `json:"requiredSkills,omitempty"`
}
