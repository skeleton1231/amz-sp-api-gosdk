/*
 * Selling Partner API for Finances
 *
 * The Selling Partner API for Finances helps you obtain financial information relevant to a seller's business. You can obtain financial events for a given order, financial event group, or date range without having to wait until a statement period closes. You can also obtain financial event groups for a given date range.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A debt payment or debt adjustment.
type DebtRecoveryEvent struct {
	// The debt recovery type.  Possible values:  * DebtPayment  * DebtPaymentFailure  *DebtAdjustment
	DebtRecoveryType string `json:"DebtRecoveryType,omitempty"`
	RecoveryAmount *Currency `json:"RecoveryAmount,omitempty"`
	OverPaymentCredit *Currency `json:"OverPaymentCredit,omitempty"`
	DebtRecoveryItemList *[]DebtRecoveryItem `json:"DebtRecoveryItemList,omitempty"`
	ChargeInstrumentList *[]ChargeInstrument `json:"ChargeInstrumentList,omitempty"`
}