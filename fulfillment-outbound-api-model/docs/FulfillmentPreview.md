# FulfillmentPreview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingSpeedCategory** | [***ShippingSpeedCategory**](ShippingSpeedCategory.md) |  | [default to null]
**ScheduledDeliveryInfo** | [***ScheduledDeliveryInfo**](ScheduledDeliveryInfo.md) |  | [optional] [default to null]
**IsFulfillable** | **bool** | When true, this fulfillment order preview is fulfillable. | [default to null]
**IsCODCapable** | **bool** | When true, this fulfillment order preview is for COD (Cash On Delivery). | [default to null]
**EstimatedShippingWeight** | [***Weight**](Weight.md) |  | [optional] [default to null]
**EstimatedFees** | [***[]Fee**](array.md) |  | [optional] [default to null]
**FulfillmentPreviewShipments** | [***[]FulfillmentPreviewShipment**](array.md) |  | [optional] [default to null]
**UnfulfillablePreviewItems** | [***[]UnfulfillablePreviewItem**](array.md) |  | [optional] [default to null]
**OrderUnfulfillableReasons** | [***[]string**](array.md) |  | [optional] [default to null]
**MarketplaceId** | **string** | The marketplace the fulfillment order is placed against. | [default to null]
**FeatureConstraints** | [**[]FeatureSettings**](FeatureSettings.md) | A list of features and their fulfillment policies to apply to the order. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

