# OrderDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrderDate** | [**time.Time**](time.Time.md) | The date the purchase order was placed. Must be in ISO-8601 date/time format. | [default to null]
**PurchaseOrderChangedDate** | [**time.Time**](time.Time.md) | The date when purchase order was last changed by Amazon after the order was placed. This date will be greater than &#x27;purchaseOrderDate&#x27;. This means the PO data was changed on that date and vendors are required to fulfill the  updated PO. The PO changes can be related to Item Quantity, Ship to Location, Ship Window etc. This field will not be present in orders that have not changed after creation. Must be in ISO-8601 date/time format. | [optional] [default to null]
**PurchaseOrderStateChangedDate** | [**time.Time**](time.Time.md) | The date when current purchase order state was changed. Current purchase order state is available in the field &#x27;purchaseOrderState&#x27;. Must be in ISO-8601 date/time format. | [default to null]
**PurchaseOrderType** | **string** | Type of purchase order. | [optional] [default to null]
**ImportDetails** | [***ImportDetails**](ImportDetails.md) |  | [optional] [default to null]
**DealCode** | **string** | If requested by the recipient, this field will contain a promotional/deal number. The discount code line is optional. It is used to obtain a price discount on items on the order. | [optional] [default to null]
**PaymentMethod** | **string** | Payment method used. | [optional] [default to null]
**BuyingParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**SellingParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**ShipToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**BillToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**ShipWindow** | **string** |  | [optional] [default to null]
**DeliveryWindow** | **string** |  | [optional] [default to null]
**Items** | [**[]OrderItem**](OrderItem.md) | A list of items in this purchase order. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

