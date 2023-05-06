# FulfillmentShipment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonShipmentId** | **string** | A shipment identifier assigned by Amazon. | [default to null]
**FulfillmentCenterId** | **string** | An identifier for the fulfillment center that the shipment will be sent from. | [default to null]
**FulfillmentShipmentStatus** | **string** | The current status of the shipment. | [default to null]
**ShippingDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**EstimatedArrivalDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**ShippingNotes** | **[]string** | Provides additional insight into shipment timeline. Primairly used to communicate that actual delivery dates aren&#x27;t available. | [optional] [default to null]
**FulfillmentShipmentItem** | [***[]FulfillmentShipmentItem**](array.md) |  | [default to null]
**FulfillmentShipmentPackage** | [***[]FulfillmentShipmentPackage**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

