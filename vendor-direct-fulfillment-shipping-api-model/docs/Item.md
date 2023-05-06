# Item

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemSequenceNumber** | **int32** | Item Sequence Number for the item. This must be the same value as sent in order for a given item. | [default to null]
**BuyerProductIdentifier** | **string** | Buyer&#x27;s Standard Identification Number (ASIN) of an item. Either buyerProductIdentifier or vendorProductIdentifier is required. | [optional] [default to null]
**VendorProductIdentifier** | **string** | The vendor selected product identification of the item. Should be the same as was sent in the purchase order, like SKU Number. | [optional] [default to null]
**ShippedQuantity** | [***ItemQuantity**](ItemQuantity.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

