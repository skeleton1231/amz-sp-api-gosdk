# OrderItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemSequenceNumber** | **string** | Numbering of the item on the purchase order. The first item will be 1, the second 2, and so on. | [default to null]
**BuyerProductIdentifier** | **string** | Buyer&#x27;s standard identification number (ASIN) of an item. | [optional] [default to null]
**VendorProductIdentifier** | **string** | The vendor selected product identification of the item. | [optional] [default to null]
**Title** | **string** | Title for the item. | [optional] [default to null]
**OrderedQuantity** | [***ItemQuantity**](ItemQuantity.md) |  | [default to null]
**ScheduledDeliveryShipment** | [***ScheduledDeliveryShipment**](ScheduledDeliveryShipment.md) |  | [optional] [default to null]
**GiftDetails** | [***GiftDetails**](GiftDetails.md) |  | [optional] [default to null]
**NetPrice** | [***Money**](Money.md) |  | [default to null]
**TaxDetails** | [***TaxItemDetails**](TaxItemDetails.md) |  | [optional] [default to null]
**TotalPrice** | [***Money**](Money.md) |  | [optional] [default to null]
**BuyerCustomizedInfo** | [***BuyerCustomizedInfoDetail**](buyerCustomizedInfoDetail.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

