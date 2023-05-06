/*
 * Selling Partner API for Finances
 *
 * The Selling Partner API for Finances helps you obtain financial information relevant to a seller's business. You can obtain financial events for a given order, financial event group, or date range without having to wait until a statement period closes. You can also obtain financial event groups for a given date range.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// A SAFE-T claim reimbursement on the seller's account.
type SafetReimbursementEvent struct {
	PostedDate *time.Time `json:"PostedDate,omitempty"`
	// A SAFE-T claim identifier.
	SAFETClaimId string `json:"SAFETClaimId,omitempty"`
	ReimbursedAmount *Currency `json:"ReimbursedAmount,omitempty"`
	// Indicates why the seller was reimbursed.
	ReasonCode string `json:"ReasonCode,omitempty"`
	SAFETReimbursementItemList *[]SafetReimbursementItem `json:"SAFETReimbursementItemList,omitempty"`
}
