# ShipmentEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonOrderId** | **string** | An Amazon-defined identifier for an order. | [optional] [default to null]
**SellerOrderId** | **string** | A seller-defined identifier for an order. | [optional] [default to null]
**MarketplaceName** | **string** | The name of the marketplace where the event occurred. | [optional] [default to null]
**OrderChargeList** | [***[]ChargeComponent**](array.md) |  | [optional] [default to null]
**OrderChargeAdjustmentList** | [***[]ChargeComponent**](array.md) |  | [optional] [default to null]
**ShipmentFeeList** | [***[]FeeComponent**](array.md) |  | [optional] [default to null]
**ShipmentFeeAdjustmentList** | [***[]FeeComponent**](array.md) |  | [optional] [default to null]
**OrderFeeList** | [***[]FeeComponent**](array.md) |  | [optional] [default to null]
**OrderFeeAdjustmentList** | [***[]FeeComponent**](array.md) |  | [optional] [default to null]
**DirectPaymentList** | [***[]DirectPayment**](array.md) |  | [optional] [default to null]
**PostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**ShipmentItemList** | [***[]ShipmentItem**](array.md) |  | [optional] [default to null]
**ShipmentItemAdjustmentList** | [***[]ShipmentItem**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

