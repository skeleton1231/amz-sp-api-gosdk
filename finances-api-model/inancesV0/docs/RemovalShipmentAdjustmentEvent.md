# RemovalShipmentAdjustmentEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**AdjustmentEventId** | **string** | The unique identifier for the adjustment event. | [optional] [default to null]
**MerchantOrderId** | **string** | The merchant removal orderId. | [optional] [default to null]
**OrderId** | **string** | The orderId for shipping inventory. | [optional] [default to null]
**TransactionType** | **string** | The type of removal order.  Possible values:  * WHOLESALE_LIQUIDATION. | [optional] [default to null]
**RemovalShipmentItemAdjustmentList** | [**[]RemovalShipmentItemAdjustment**](RemovalShipmentItemAdjustment.md) | A comma-delimited list of Removal shipmentItemAdjustment details for FBA inventory. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

