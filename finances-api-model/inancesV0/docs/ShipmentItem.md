# ShipmentItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerSKU** | **string** | The seller SKU of the item. The seller SKU is qualified by the seller&#x27;s seller ID, which is included with every call to the Selling Partner API. | [optional] [default to null]
**OrderItemId** | **string** | An Amazon-defined order item identifier. | [optional] [default to null]
**OrderAdjustmentItemId** | **string** | An Amazon-defined order adjustment identifier defined for refunds, guarantee claims, and chargeback events. | [optional] [default to null]
**QuantityShipped** | **int32** | The number of items shipped. | [optional] [default to null]
**ItemChargeList** | [***[]ChargeComponent**](array.md) |  | [optional] [default to null]
**ItemChargeAdjustmentList** | [***[]ChargeComponent**](array.md) |  | [optional] [default to null]
**ItemFeeList** | [***[]FeeComponent**](array.md) |  | [optional] [default to null]
**ItemFeeAdjustmentList** | [***[]FeeComponent**](array.md) |  | [optional] [default to null]
**ItemTaxWithheldList** | [***[]TaxWithheldComponent**](array.md) |  | [optional] [default to null]
**PromotionList** | [***[]Promotion**](array.md) |  | [optional] [default to null]
**PromotionAdjustmentList** | [***[]Promotion**](array.md) |  | [optional] [default to null]
**CostOfPointsGranted** | [***Currency**](Currency.md) |  | [optional] [default to null]
**CostOfPointsReturned** | [***Currency**](Currency.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

