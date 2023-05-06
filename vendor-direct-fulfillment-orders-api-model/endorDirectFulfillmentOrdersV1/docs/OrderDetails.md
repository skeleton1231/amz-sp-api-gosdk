# OrderDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerOrderNumber** | **string** | The customer order number. | [default to null]
**OrderDate** | [**time.Time**](time.Time.md) | The date the order was placed. This field is expected to be in ISO-8601 date/time format, for example:2018-07-16T23:00:00Z/ 2018-07-16T23:00:00-05:00 /2018-07-16T23:00:00-08:00. If no time zone is specified, UTC should be assumed. | [default to null]
**OrderStatus** | **string** | Current status of the order. | [optional] [default to null]
**ShipmentDetails** | [***ShipmentDetails**](ShipmentDetails.md) |  | [default to null]
**TaxTotal** | [***OrderDetailsTaxTotal**](OrderDetails_taxTotal.md) |  | [optional] [default to null]
**SellingParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**ShipFromParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**ShipToParty** | [***Address**](Address.md) |  | [default to null]
**BillToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**Items** | [**[]OrderItem**](OrderItem.md) | A list of items in this purchase order. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

