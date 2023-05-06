# RetrochargeEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetrochargeEventType** | **string** | The type of event.  Possible values:  * Retrocharge  * RetrochargeReversal | [optional] [default to null]
**AmazonOrderId** | **string** | An Amazon-defined identifier for an order. | [optional] [default to null]
**PostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**BaseTax** | [***Currency**](Currency.md) |  | [optional] [default to null]
**ShippingTax** | [***Currency**](Currency.md) |  | [optional] [default to null]
**MarketplaceName** | **string** | The name of the marketplace where the retrocharge event occurred. | [optional] [default to null]
**RetrochargeTaxWithheldList** | [***[]TaxWithheldComponent**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

