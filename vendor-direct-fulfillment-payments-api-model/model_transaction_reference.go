/*
 * Selling Partner API for Direct Fulfillment Payments
 *
 * The Selling Partner API for Direct Fulfillment Payments provides programmatic access to a direct fulfillment vendor's invoice data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TransactionReference struct {
	// GUID to identify this transaction. This value can be used with the Transaction Status API to return the status of this transaction.
	TransactionId string `json:"transactionId,omitempty"`
}
