# InvoiceDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNumber** | **string** | The unique invoice number. | [default to null]
**InvoiceDate** | [**time.Time**](time.Time.md) | Invoice date. | [default to null]
**ReferenceNumber** | **string** | An additional unique reference number used for regulatory or other purposes. | [optional] [default to null]
**RemitToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**ShipFromParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**BillToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**ShipToCountryCode** | **string** | Ship-to country code. | [optional] [default to null]
**PaymentTermsCode** | **string** | The payment terms for the invoice. | [optional] [default to null]
**InvoiceTotal** | [***Money**](Money.md) |  | [default to null]
**TaxTotals** | [**[]TaxDetail**](TaxDetail.md) | Individual tax details per line item. | [optional] [default to null]
**AdditionalDetails** | [**[]AdditionalDetails**](AdditionalDetails.md) | Additional details provided by the selling party, for tax-related or other purposes. | [optional] [default to null]
**ChargeDetails** | [**[]ChargeDetails**](ChargeDetails.md) | Total charge amount details for all line items. | [optional] [default to null]
**Items** | [**[]InvoiceItem**](InvoiceItem.md) | Provides the details of the items in this invoice. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

