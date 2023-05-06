# OrderItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemSequenceNumber** | **string** | Numbering of the item on the purchase order. The first item will be 1, the second 2, and so on. | [default to null]
**AmazonProductIdentifier** | **string** | Amazon Standard Identification Number (ASIN) of an item. | [optional] [default to null]
**VendorProductIdentifier** | **string** | The vendor selected product identification of the item. | [optional] [default to null]
**OrderedQuantity** | [***ItemQuantity**](ItemQuantity.md) |  | [default to null]
**IsBackOrderAllowed** | **bool** | When true, we will accept backorder confirmations for this item. | [default to null]
**NetCost** | [***Money**](Money.md) |  | [optional] [default to null]
**ListPrice** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

