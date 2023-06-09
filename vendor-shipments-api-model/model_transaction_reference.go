/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TransactionReference struct {
	// GUID assigned by Buyer to identify this transaction. This value can be used with the Transaction Status API to return the status of this transaction.
	TransactionId string `json:"transactionId,omitempty"`
}
