# AffordabilityExpenseEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonOrderId** | **string** | An Amazon-defined identifier for an order. | [optional] [default to null]
**PostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**MarketplaceId** | **string** | An encrypted, Amazon-defined marketplace identifier. | [optional] [default to null]
**TransactionType** | **string** | Indicates the type of transaction.   Possible values:  * Charge - For an affordability promotion expense.  * Refund - For an affordability promotion expense reversal. | [optional] [default to null]
**BaseExpense** | [***Currency**](Currency.md) |  | [optional] [default to null]
**TaxTypeCGST** | [***Currency**](Currency.md) |  | [default to null]
**TaxTypeSGST** | [***Currency**](Currency.md) |  | [default to null]
**TaxTypeIGST** | [***Currency**](Currency.md) |  | [default to null]
**TotalExpense** | [***Currency**](Currency.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

