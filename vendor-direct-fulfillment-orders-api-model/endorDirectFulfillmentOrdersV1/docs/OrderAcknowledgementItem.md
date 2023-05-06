# OrderAcknowledgementItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrderNumber** | **string** | The purchase order number for this order. Formatting Notes: alpha-numeric code. | [default to null]
**VendorOrderNumber** | **string** | The vendor&#x27;s order number for this order. | [default to null]
**AcknowledgementDate** | [**time.Time**](time.Time.md) | The date and time when the order is acknowledged, in ISO-8601 date/time format. For example: 2018-07-16T23:00:00Z / 2018-07-16T23:00:00-05:00 / 2018-07-16T23:00:00-08:00. | [default to null]
**AcknowledgementStatus** | [***AcknowledgementStatus**](AcknowledgementStatus.md) |  | [default to null]
**SellingParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**ShipFromParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**ItemAcknowledgements** | [**[]OrderItemAcknowledgement**](OrderItemAcknowledgement.md) | Item details including acknowledged quantity. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

