# Containers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerType** | **string** | The type of container. | [default to null]
**ContainerSequenceNumber** | **string** | An integer that must be submitted for multi-box shipments only, where one item may come in separate packages. | [optional] [default to null]
**ContainerIdentifiers** | [**[]ContainerIdentification**](ContainerIdentification.md) | A list of carton identifiers. | [default to null]
**TrackingNumber** | **string** | The tracking number used for identifying the shipment. | [optional] [default to null]
**Dimensions** | [***Dimensions**](Dimensions.md) |  | [optional] [default to null]
**Weight** | [***Weight**](Weight.md) |  | [optional] [default to null]
**Tier** | **int32** | Number of layers per pallet. | [optional] [default to null]
**Block** | **int32** | Number of cartons per layer on the pallet. | [optional] [default to null]
**InnerContainersDetails** | [***InnerContainersDetails**](InnerContainersDetails.md) |  | [optional] [default to null]
**PackedItems** | [**[]PackedItems**](PackedItems.md) | A list of packed items. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

