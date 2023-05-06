# SolutionProviderCreditEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderTransactionType** | **string** | The transaction type. | [optional] [default to null]
**SellerOrderId** | **string** | A seller-defined identifier for an order. | [optional] [default to null]
**MarketplaceId** | **string** | The identifier of the marketplace where the order was placed. | [optional] [default to null]
**MarketplaceCountryCode** | **string** | The two-letter country code of the country associated with the marketplace where the order was placed. | [optional] [default to null]
**SellerId** | **string** | The Amazon-defined identifier of the seller. | [optional] [default to null]
**SellerStoreName** | **string** | The store name where the payment event occurred. | [optional] [default to null]
**ProviderId** | **string** | The Amazon-defined identifier of the solution provider. | [optional] [default to null]
**ProviderStoreName** | **string** | The store name where the payment event occurred. | [optional] [default to null]
**TransactionAmount** | [***Currency**](Currency.md) |  | [optional] [default to null]
**TransactionCreationDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

