# CreateFulfillmentOrderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | The marketplace the fulfillment order is placed against. | [optional] [default to null]
**SellerFulfillmentOrderId** | **string** | A fulfillment order identifier that the seller creates to track their fulfillment order. The SellerFulfillmentOrderId must be unique for each fulfillment order that a seller creates. If the seller&#x27;s system already creates unique order identifiers, then these might be good values for them to use. | [default to null]
**DisplayableOrderId** | **string** | A fulfillment order identifier that the seller creates. This value displays as the order identifier in recipient-facing materials such as the outbound shipment packing slip. The value of DisplayableOrderId should match the order identifier that the seller provides to the recipient. The seller can use the SellerFulfillmentOrderId for this value or they can specify an alternate value if they want the recipient to reference an alternate order identifier.  The value must be an alpha-numeric or ISO 8859-1 compliant string from one to 40 characters in length. Cannot contain two spaces in a row. Leading and trailing white space is removed. | [default to null]
**DisplayableOrderDate** | [***time.Time**](time.Time.md) |  | [default to null]
**DisplayableOrderComment** | **string** | Order-specific text that appears in recipient-facing materials such as the outbound shipment packing slip. | [default to null]
**ShippingSpeedCategory** | [***ShippingSpeedCategory**](ShippingSpeedCategory.md) |  | [default to null]
**DeliveryWindow** | [***DeliveryWindow**](DeliveryWindow.md) |  | [optional] [default to null]
**DestinationAddress** | [***Address**](Address.md) |  | [default to null]
**FulfillmentAction** | [***FulfillmentAction**](FulfillmentAction.md) |  | [optional] [default to null]
**FulfillmentPolicy** | [***FulfillmentPolicy**](FulfillmentPolicy.md) |  | [optional] [default to null]
**CodSettings** | [***CodSettings**](CODSettings.md) |  | [optional] [default to null]
**ShipFromCountryCode** | **string** | The two-character country code for the country from which the fulfillment order ships. Must be in ISO 3166-1 alpha-2 format. | [optional] [default to null]
**NotificationEmails** | [***[]string**](array.md) |  | [optional] [default to null]
**FeatureConstraints** | [**[]FeatureSettings**](FeatureSettings.md) | A list of features and their fulfillment policies to apply to the order. | [optional] [default to null]
**Items** | [***[]CreateFulfillmentOrderItem**](array.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

