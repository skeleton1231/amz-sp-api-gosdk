# GetFulfillmentPreviewRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | The marketplace the fulfillment order is placed against. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [default to null]
**Items** | [***[]GetFulfillmentPreviewItem**](array.md) |  | [default to null]
**ShippingSpeedCategories** | [***[]ShippingSpeedCategory**](array.md) |  | [optional] [default to null]
**IncludeCODFulfillmentPreview** | **bool** | When true, returns all fulfillment order previews both for COD and not for COD. Otherwise, returns only fulfillment order previews that are not for COD. | [optional] [default to null]
**IncludeDeliveryWindows** | **bool** | When true, returns the ScheduledDeliveryInfo response object, which contains the available delivery windows for a Scheduled Delivery. The ScheduledDeliveryInfo response object can only be returned for fulfillment order previews with ShippingSpeedCategories &#x3D; ScheduledDelivery. | [optional] [default to null]
**FeatureConstraints** | [**[]FeatureSettings**](FeatureSettings.md) | A list of features and their fulfillment policies to apply to the order. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

