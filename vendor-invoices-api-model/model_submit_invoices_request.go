/*
 * Selling Partner API for Retail Procurement Payments
 *
 * The Selling Partner API for Retail Procurement Payments provides programmatic access to vendors payments data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The request schema for the submitInvoices operation.
type SubmitInvoicesRequest struct {
	Invoices []Invoice `json:"invoices,omitempty"`
}
