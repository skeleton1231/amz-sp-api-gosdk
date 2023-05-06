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

// A fee event related to Amazon Imaging services.
type ImagingServicesFeeEvent struct {
	// The identifier for the imaging services request.
	ImagingRequestBillingItemID string `json:"ImagingRequestBillingItemID,omitempty"`
	// The Amazon Standard Identification Number (ASIN) of the item for which the imaging service was requested.
	ASIN string `json:"ASIN,omitempty"`
	PostedDate *time.Time `json:"PostedDate,omitempty"`
	FeeList *[]FeeComponent `json:"FeeList,omitempty"`
}