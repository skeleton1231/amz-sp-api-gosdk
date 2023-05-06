# NetworkComminglingTransactionEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionType** | **string** | The type of network item swap.  Possible values:  * NetCo - A Fulfillment by Amazon inventory pooling transaction. Available only in the India marketplace.  * ComminglingVAT - A commingling VAT transaction. Available only in the UK, Spain, France, Germany, and Italy marketplaces. | [optional] [default to null]
**PostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**NetCoTransactionID** | **string** | The identifier for the network item swap. | [optional] [default to null]
**SwapReason** | **string** | The reason for the network item swap. | [optional] [default to null]
**ASIN** | **string** | The Amazon Standard Identification Number (ASIN) of the swapped item. | [optional] [default to null]
**MarketplaceId** | **string** | The marketplace in which the event took place. | [optional] [default to null]
**TaxExclusiveAmount** | [***Currency**](Currency.md) |  | [optional] [default to null]
**TaxAmount** | [***Currency**](Currency.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

