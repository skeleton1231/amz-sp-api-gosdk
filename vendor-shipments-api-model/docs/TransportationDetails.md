# TransportationDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipMode** | **string** | The type of shipment. | [optional] [default to null]
**TransportationMode** | **string** | The mode of transportation for this shipment. | [optional] [default to null]
**ShippedDate** | [**time.Time**](time.Time.md) | Date when shipment is performed by the Vendor to Buyer | [optional] [default to null]
**EstimatedDeliveryDate** | [**time.Time**](time.Time.md) | Estimated Date on which shipment will be delivered from Vendor to Buyer | [optional] [default to null]
**ShipmentDeliveryDate** | [**time.Time**](time.Time.md) | Date on which shipment will be delivered from Vendor to Buyer | [optional] [default to null]
**CarrierDetails** | [***CarrierDetails**](CarrierDetails.md) |  | [optional] [default to null]
**BillOfLadingNumber** | **string** | Bill Of Lading (BOL) number is the unique number assigned by the vendor. The BOL present in the Shipment Confirmation message ideally matches the paper BOL provided with the shipment, but that is no must. Instead of BOL, an alternative reference number (like Delivery Note Number) for the shipment can also be sent in this field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

