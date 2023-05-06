# RemovalShipmentEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**MerchantOrderId** | **string** | The merchant removal orderId. | [optional] [default to null]
**OrderId** | **string** | The identifier for the removal shipment order. | [optional] [default to null]
**TransactionType** | **string** | The type of removal order.  Possible values:  * WHOLESALE_LIQUIDATION | [optional] [default to null]
**RemovalShipmentItemList** | [***[]RemovalShipmentItem**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

