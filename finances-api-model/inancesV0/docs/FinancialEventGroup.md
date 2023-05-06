# FinancialEventGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinancialEventGroupId** | **string** | A unique identifier for the financial event group. | [optional] [default to null]
**ProcessingStatus** | **string** | The processing status of the financial event group indicates whether the balance of the financial event group is settled.  Possible values:  * Open  * Closed | [optional] [default to null]
**FundTransferStatus** | **string** | The status of the fund transfer. | [optional] [default to null]
**OriginalTotal** | [***Currency**](Currency.md) |  | [optional] [default to null]
**ConvertedTotal** | [***Currency**](Currency.md) |  | [optional] [default to null]
**FundTransferDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**TraceId** | **string** | The trace identifier used by sellers to look up transactions externally. | [optional] [default to null]
**AccountTail** | **string** | The account tail of the payment instrument. | [optional] [default to null]
**BeginningBalance** | [***Currency**](Currency.md) |  | [optional] [default to null]
**FinancialEventGroupStart** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**FinancialEventGroupEnd** | [***time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

