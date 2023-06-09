/*
 * Selling Partner API for Retail Procurement Payments
 *
 * The Selling Partner API for Retail Procurement Payments provides programmatic access to vendors payments data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response schema for the submitInvoices operation.
type SubmitInvoicesResponse struct {
	Payload *TransactionId `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
