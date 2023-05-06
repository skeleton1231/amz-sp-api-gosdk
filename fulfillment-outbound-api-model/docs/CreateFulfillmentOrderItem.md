# CreateFulfillmentOrderItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerSku** | **string** | The seller SKU of the item. | [default to null]
**SellerFulfillmentOrderItemId** | **string** | A fulfillment order item identifier that the seller creates to track fulfillment order items. Used to disambiguate multiple fulfillment items that have the same SellerSKU. For example, the seller might assign different SellerFulfillmentOrderItemId values to two items in a fulfillment order that share the same SellerSKU but have different GiftMessage values. | [default to null]
**Quantity** | **int32** |  | [default to null]
**GiftMessage** | **string** | A message to the gift recipient, if applicable. | [optional] [default to null]
**DisplayableComment** | **string** | Item-specific text that displays in recipient-facing materials such as the outbound shipment packing slip. | [optional] [default to null]
**FulfillmentNetworkSku** | **string** | Amazon&#x27;s fulfillment network SKU of the item. | [optional] [default to null]
**PerUnitDeclaredValue** | [***Money**](Money.md) |  | [optional] [default to null]
**PerUnitPrice** | [***Money**](Money.md) |  | [optional] [default to null]
**PerUnitTax** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

