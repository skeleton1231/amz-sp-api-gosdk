# OrderAcknowledgementItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemSequenceNumber** | **string** | Line item sequence number for the item. | [optional] [default to null]
**AmazonProductIdentifier** | **string** | Amazon Standard Identification Number (ASIN) of an item. | [optional] [default to null]
**VendorProductIdentifier** | **string** | The vendor selected product identification of the item. Should be the same as was sent in the purchase order. | [optional] [default to null]
**OrderedQuantity** | [***ItemQuantity**](ItemQuantity.md) |  | [default to null]
**NetCost** | [***Money**](Money.md) |  | [optional] [default to null]
**ListPrice** | [***Money**](Money.md) |  | [optional] [default to null]
**DiscountMultiplier** | **string** | The discount multiplier that should be applied to the price if a vendor sells books with a list price. This is a multiplier factor to arrive at a final discounted price. A multiplier of .90 would be the factor if a 10% discount is given. | [optional] [default to null]
**ItemAcknowledgements** | [**[]OrderItemAcknowledgement**](OrderItemAcknowledgement.md) | This is used to indicate acknowledged quantity. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

