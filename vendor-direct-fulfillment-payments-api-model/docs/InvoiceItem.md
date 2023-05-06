# InvoiceItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemSequenceNumber** | **string** | Numbering of the item on the purchase order. The first item will be 1, the second 2, and so on. | [default to null]
**BuyerProductIdentifier** | **string** | Buyer&#x27;s standard identification number (ASIN) of an item. | [optional] [default to null]
**VendorProductIdentifier** | **string** | The vendor selected product identification of the item. | [optional] [default to null]
**InvoicedQuantity** | [***ItemQuantity**](ItemQuantity.md) |  | [default to null]
**NetCost** | [***Money**](Money.md) |  | [default to null]
**PurchaseOrderNumber** | **string** | The purchase order number for this order. Formatting Notes: 8-character alpha-numeric code. | [default to null]
**VendorOrderNumber** | **string** | The vendor&#x27;s order number for this order. | [optional] [default to null]
**HsnCode** | **string** | Harmonized System of Nomenclature (HSN) tax code. The HSN number cannot contain alphabets. | [optional] [default to null]
**TaxDetails** | [**[]TaxDetail**](TaxDetail.md) | Individual tax details per line item. | [optional] [default to null]
**ChargeDetails** | [**[]ChargeDetails**](ChargeDetails.md) | Individual charge details per line item. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

