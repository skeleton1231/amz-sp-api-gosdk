# OrderStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrderNumber** | **string** | The buyer&#x27;s purchase order number for this order. Formatting Notes: 8-character alpha-numeric code. | [default to null]
**PurchaseOrderStatus** | **string** | The status of the buyer&#x27;s purchase order for this order. | [default to null]
**PurchaseOrderDate** | [**time.Time**](time.Time.md) | The date the purchase order was placed. Must be in ISO-8601 date/time format. | [default to null]
**LastUpdatedDate** | [**time.Time**](time.Time.md) | The date when the purchase order was last updated. Must be in ISO-8601 date/time format. | [optional] [default to null]
**SellingParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**ShipToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**ItemStatus** | [***[]OrderItemStatus**](array.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

