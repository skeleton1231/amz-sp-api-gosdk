# InboundShipmentInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentId** | **string** | The shipment identifier submitted in the request. | [optional] [default to null]
**ShipmentName** | **string** | The name for the inbound shipment. | [optional] [default to null]
**ShipFromAddress** | [***Address**](Address.md) |  | [default to null]
**DestinationFulfillmentCenterId** | **string** | An Amazon fulfillment center identifier created by Amazon. | [optional] [default to null]
**ShipmentStatus** | [***ShipmentStatus**](ShipmentStatus.md) |  | [optional] [default to null]
**LabelPrepType** | [***LabelPrepType**](LabelPrepType.md) |  | [optional] [default to null]
**AreCasesRequired** | **bool** | Indicates whether or not an inbound shipment contains case-packed boxes. When AreCasesRequired &#x3D; true for an inbound shipment, all items in the inbound shipment must be case packed. | [default to null]
**ConfirmedNeedByDate** | **string** |  | [optional] [default to null]
**BoxContentsSource** | [***BoxContentsSource**](BoxContentsSource.md) |  | [optional] [default to null]
**EstimatedBoxContentsFee** | [***BoxContentsFeeDetails**](BoxContentsFeeDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

