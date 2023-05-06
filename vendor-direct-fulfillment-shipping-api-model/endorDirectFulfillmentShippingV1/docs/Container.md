# Container

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerType** | **string** | The type of container. | [default to null]
**ContainerIdentifier** | **string** | The container identifier. | [default to null]
**TrackingNumber** | **string** | The tracking number. | [optional] [default to null]
**ManifestId** | **string** | The manifest identifier. | [optional] [default to null]
**ManifestDate** | **string** | The date of the manifest. | [optional] [default to null]
**ShipMethod** | **string** | The shipment method. | [optional] [default to null]
**ScacCode** | **string** | SCAC code required for NA VOC vendors only. | [optional] [default to null]
**Carrier** | **string** | Carrier required for EU VOC vendors only. | [optional] [default to null]
**ContainerSequenceNumber** | **int32** | An integer that must be submitted for multi-box shipments only, where one item may come in separate packages. | [optional] [default to null]
**Dimensions** | [***Dimensions**](Dimensions.md) |  | [optional] [default to null]
**Weight** | [***Weight**](Weight.md) |  | [optional] [default to null]
**PackedItems** | [**[]PackedItem**](PackedItem.md) | A list of packed items. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

