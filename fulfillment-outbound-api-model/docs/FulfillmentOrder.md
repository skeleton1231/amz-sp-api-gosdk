# FulfillmentOrder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerFulfillmentOrderId** | **string** | The fulfillment order identifier submitted with the createFulfillmentOrder operation. | [default to null]
**MarketplaceId** | **string** | The identifier for the marketplace the fulfillment order is placed against. | [default to null]
**DisplayableOrderId** | **string** | A fulfillment order identifier submitted with the createFulfillmentOrder operation. Displays as the order identifier in recipient-facing materials such as the packing slip. | [default to null]
**DisplayableOrderDate** | [***time.Time**](time.Time.md) |  | [default to null]
**DisplayableOrderComment** | **string** | A text block submitted with the createFulfillmentOrder operation. Displays in recipient-facing materials such as the packing slip. | [default to null]
**ShippingSpeedCategory** | [***ShippingSpeedCategory**](ShippingSpeedCategory.md) |  | [default to null]
**DeliveryWindow** | [***DeliveryWindow**](DeliveryWindow.md) |  | [optional] [default to null]
**DestinationAddress** | [***Address**](Address.md) |  | [default to null]
**FulfillmentAction** | [***FulfillmentAction**](FulfillmentAction.md) |  | [optional] [default to null]
**FulfillmentPolicy** | [***FulfillmentPolicy**](FulfillmentPolicy.md) |  | [optional] [default to null]
**CodSettings** | [***CodSettings**](CODSettings.md) |  | [optional] [default to null]
**ReceivedDate** | [***time.Time**](time.Time.md) |  | [default to null]
**FulfillmentOrderStatus** | [***FulfillmentOrderStatus**](FulfillmentOrderStatus.md) |  | [default to null]
**StatusUpdatedDate** | [***time.Time**](time.Time.md) |  | [default to null]
**NotificationEmails** | [***[]string**](array.md) |  | [optional] [default to null]
**FeatureConstraints** | [**[]FeatureSettings**](FeatureSettings.md) | A list of features and their fulfillment policies to apply to the order. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

