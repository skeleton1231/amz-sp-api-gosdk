# ShipmentInformation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorDetails** | [***VendorDetails**](VendorDetails.md) |  | [optional] [default to null]
**BuyerReferenceNumber** | **string** | Buyer Reference number which is a unique number. | [optional] [default to null]
**ShipToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**ShipFromParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**WarehouseId** | **string** | Vendor Warehouse ID from where the shipment is scheduled to be picked up by buyer / Carrier. | [optional] [default to null]
**MasterTrackingId** | **string** | Unique Id with  which  the shipment can be tracked for Small Parcels. | [optional] [default to null]
**TotalLabelCount** | **int32** | Number of Labels that are created as part of this shipment. | [optional] [default to null]
**ShipMode** | **string** | Type of shipment whether it is Small Parcel | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

