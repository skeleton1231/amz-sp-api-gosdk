# FulfillmentOrderItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerSku** | **string** | The seller SKU of the item. | [default to null]
**SellerFulfillmentOrderItemId** | **string** | A fulfillment order item identifier submitted with a call to the createFulfillmentOrder operation. | [default to null]
**Quantity** | **int32** |  | [default to null]
**GiftMessage** | **string** | A message to the gift recipient, if applicable. | [optional] [default to null]
**DisplayableComment** | **string** | Item-specific text that displays in recipient-facing materials such as the outbound shipment packing slip. | [optional] [default to null]
**FulfillmentNetworkSku** | **string** | Amazon&#x27;s fulfillment network SKU of the item. | [optional] [default to null]
**OrderItemDisposition** | **string** | Indicates whether the item is sellable or unsellable. | [optional] [default to null]
**CancelledQuantity** | **int32** |  | [default to null]
**UnfulfillableQuantity** | **int32** |  | [default to null]
**EstimatedShipDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**EstimatedArrivalDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**PerUnitPrice** | [***Money**](Money.md) |  | [optional] [default to null]
**PerUnitTax** | [***Money**](Money.md) |  | [optional] [default to null]
**PerUnitDeclaredValue** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

