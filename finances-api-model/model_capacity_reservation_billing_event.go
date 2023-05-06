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

// An event related to a capacity reservation billing charge.
type CapacityReservationBillingEvent struct {
	// Indicates the type of transaction. For example, FBA Inventory Fee
	TransactionType string `json:"TransactionType,omitempty"`
	PostedDate *time.Time `json:"PostedDate,omitempty"`
	// A short description of the capacity reservation billing event.
	Description string `json:"Description,omitempty"`
	TransactionAmount *Currency `json:"TransactionAmount,omitempty"`
}
