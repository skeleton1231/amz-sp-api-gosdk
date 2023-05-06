# InvoiceItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemSequenceNumber** | **int32** | Unique number related to this line item. | [default to null]
**AmazonProductIdentifier** | **string** | Amazon Standard Identification Number (ASIN) of an item. | [optional] [default to null]
**VendorProductIdentifier** | **string** | The vendor selected product identifier of the item. Should be the same as was provided in the purchase order. | [optional] [default to null]
**InvoicedQuantity** | [***ItemQuantity**](ItemQuantity.md) |  | [default to null]
**NetCost** | [***Money**](Money.md) |  | [default to null]
**PurchaseOrderNumber** | **string** | The Amazon purchase order number for this invoiced line item. Formatting Notes: 8-character alpha-numeric code. This value is mandatory only when invoiceType is Invoice, and is not required when invoiceType is CreditNote. | [optional] [default to null]
**HsnCode** | **string** | HSN Tax code. The HSN number cannot contain alphabets. | [optional] [default to null]
**CreditNoteDetails** | [***CreditNoteDetails**](CreditNoteDetails.md) |  | [optional] [default to null]
**TaxDetails** | [**[]TaxDetails**](TaxDetails.md) | Individual tax details per line item. | [optional] [default to null]
**ChargeDetails** | [**[]ChargeDetails**](ChargeDetails.md) | Individual charge details per line item. | [optional] [default to null]
**AllowanceDetails** | [**[]AllowanceDetails**](AllowanceDetails.md) | Individual allowance details per line item. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

