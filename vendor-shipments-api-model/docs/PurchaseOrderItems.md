# PurchaseOrderItems

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemSequenceNumber** | **string** | Item sequence number for the item. The first item will be 001, the second 002, and so on. This number is used as a reference to refer to this item from the carton or pallet level. | [default to null]
**BuyerProductIdentifier** | **string** | Amazon Standard Identification Number (ASIN) for a SKU | [optional] [default to null]
**VendorProductIdentifier** | **string** | The vendor selected product identification of the item. Should be the same as was sent in the purchase order. | [optional] [default to null]
**ShippedQuantity** | [***ItemQuantity**](ItemQuantity.md) |  | [default to null]
**MaximumRetailPrice** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

