/*
 * Selling Partner API for Direct Fulfillment Orders
 *
 * The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Pagination struct {
	// A generated string used to pass information to your next request. If NextToken is returned, pass the value of NextToken to the next request. If NextToken is not returned, there are no more order items to return.
	NextToken string `json:"nextToken,omitempty"`
}
