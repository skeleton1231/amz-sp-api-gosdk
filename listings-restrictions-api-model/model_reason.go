/*
 * Selling Partner API for Listings Restrictions
 *
 * The Selling Partner API for Listings Restrictions provides programmatic access to restrictions on Amazon catalog listings.  For more information, see the [Listings Restrictions API Use Case Guide](doc:listings-restrictions-api-v2021-08-01-use-case-guide).
 *
 * API version: 2021-08-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A reason for the restriction, including path forward links that may allow Selling Partners to remove the restriction, if available.
type Reason struct {
	// A message describing the reason for the restriction.
	Message string `json:"message"`
	// A code indicating why the listing is restricted.
	ReasonCode string `json:"reasonCode,omitempty"`
	// A list of path forward links that may allow Selling Partners to remove the restriction.
	Links []Link `json:"links,omitempty"`
}
