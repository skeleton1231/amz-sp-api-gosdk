# OrderAcknowledgement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrderNumber** | **string** | The purchase order number. Formatting Notes: 8-character alpha-numeric code. | [default to null]
**SellingParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**AcknowledgementDate** | [**time.Time**](time.Time.md) | The date and time when the purchase order is acknowledged, in ISO-8601 date/time format. | [default to null]
**Items** | [**[]OrderAcknowledgementItem**](OrderAcknowledgementItem.md) | A list of the items being acknowledged with associated details. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

