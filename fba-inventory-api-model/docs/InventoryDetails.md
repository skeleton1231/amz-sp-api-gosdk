# InventoryDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillableQuantity** | **int32** | The item quantity that can be picked, packed, and shipped. | [optional] [default to null]
**InboundWorkingQuantity** | **int32** | The number of units in an inbound shipment for which you have notified Amazon. | [optional] [default to null]
**InboundShippedQuantity** | **int32** | The number of units in an inbound shipment that you have notified Amazon about and have provided a tracking number. | [optional] [default to null]
**InboundReceivingQuantity** | **int32** | The number of units that have not yet been received at an Amazon fulfillment center for processing, but are part of an inbound shipment with some units that have already been received and processed. | [optional] [default to null]
**ReservedQuantity** | [***ReservedQuantity**](ReservedQuantity.md) |  | [optional] [default to null]
**ResearchingQuantity** | [***ResearchingQuantity**](ResearchingQuantity.md) |  | [optional] [default to null]
**UnfulfillableQuantity** | [***UnfulfillableQuantity**](UnfulfillableQuantity.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

