# ProductAdsPaymentEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**TransactionType** | **string** | Indicates if the transaction is for a charge or a refund.  Possible values:  * charge - Charge  * refund - Refund | [optional] [default to null]
**InvoiceId** | **string** | Identifier for the invoice that the transaction appears in. | [optional] [default to null]
**BaseValue** | [***Currency**](Currency.md) |  | [optional] [default to null]
**TaxValue** | [***Currency**](Currency.md) |  | [optional] [default to null]
**TransactionValue** | [***Currency**](Currency.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

