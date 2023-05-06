# OrderItemStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemSequenceNumber** | **string** | Numbering of the item on the purchase order. The first item will be 1, the second 2, and so on. | [default to null]
**BuyerProductIdentifier** | **string** | Buyer&#x27;s Standard Identification Number (ASIN) of an item. | [optional] [default to null]
**VendorProductIdentifier** | **string** | The vendor selected product identification of the item. | [optional] [default to null]
**NetCost** | [***Money**](Money.md) |  | [optional] [default to null]
**ListPrice** | [***Money**](Money.md) |  | [optional] [default to null]
**OrderedQuantity** | [***OrderItemStatusOrderedQuantity**](OrderItemStatus_orderedQuantity.md) |  | [optional] [default to null]
**AcknowledgementStatus** | [***OrderItemStatusAcknowledgementStatus**](OrderItemStatus_acknowledgementStatus.md) |  | [optional] [default to null]
**ReceivingStatus** | [***OrderItemStatusReceivingStatus**](OrderItemStatus_receivingStatus.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

