# ShipmentConfirmation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrderNumber** | **string** | Purchase order number corresponding to the shipment. | [default to null]
**ShipmentDetails** | [***ShipmentDetails**](ShipmentDetails.md) |  | [default to null]
**SellingParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**ShipFromParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**Items** | [**[]Item**](Item.md) | Provide the details of the items in this shipment. If any of the item details field is common at a package or a pallet level, then provide them at the corresponding package. | [default to null]
**Containers** | [**[]Container**](Container.md) | Provide the details of the items in this shipment. If any of the item details field is common at a package or a pallet level, then provide them at the corresponding package. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

