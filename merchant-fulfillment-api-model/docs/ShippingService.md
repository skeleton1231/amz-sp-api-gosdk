# ShippingService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingServiceName** | **string** | A plain text representation of a carrier&#x27;s shipping service. For example, \&quot;UPS Ground\&quot; or \&quot;FedEx Standard Overnight\&quot;.  | [default to null]
**CarrierName** | **string** | The name of the carrier. | [default to null]
**ShippingServiceId** | **string** |  | [default to null]
**ShippingServiceOfferId** | **string** | An Amazon-defined shipping service offer identifier. | [default to null]
**ShipDate** | [***time.Time**](time.Time.md) |  | [default to null]
**EarliestEstimatedDeliveryDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**LatestEstimatedDeliveryDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**Rate** | [***CurrencyAmount**](CurrencyAmount.md) |  | [default to null]
**ShippingServiceOptions** | [***ShippingServiceOptions**](ShippingServiceOptions.md) |  | [default to null]
**AvailableShippingServiceOptions** | [***AvailableShippingServiceOptions**](AvailableShippingServiceOptions.md) |  | [optional] [default to null]
**AvailableLabelFormats** | [***[]LabelFormat**](array.md) |  | [optional] [default to null]
**AvailableFormatOptionsForLabel** | [***[]LabelFormatOption**](array.md) |  | [optional] [default to null]
**RequiresAdditionalSellerInputs** | **bool** | When true, additional seller inputs are required. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

