# InboundShipmentPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentId** | **string** | A shipment identifier originally returned by the createInboundShipmentPlan operation. | [default to null]
**DestinationFulfillmentCenterId** | **string** | An Amazon fulfillment center identifier created by Amazon. | [default to null]
**ShipToAddress** | [***Address**](Address.md) |  | [default to null]
**LabelPrepType** | [***LabelPrepType**](LabelPrepType.md) |  | [default to null]
**Items** | [***[]InboundShipmentPlanItem**](array.md) |  | [default to null]
**EstimatedBoxContentsFee** | [***BoxContentsFeeDetails**](BoxContentsFeeDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

