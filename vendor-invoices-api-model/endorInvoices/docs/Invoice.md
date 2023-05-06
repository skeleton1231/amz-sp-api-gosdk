# Invoice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceType** | **string** | Identifies the type of invoice. | [default to null]
**Id** | **string** | Unique number relating to the charges defined in this document. This will be invoice number if the document type is Invoice or CreditNote number if the document type is Credit Note. Failure to provide this reference will result in a rejection. | [default to null]
**ReferenceNumber** | **string** | An additional unique reference number used for regulatory or other purposes. | [optional] [default to null]
**Date** | [***time.Time**](time.Time.md) |  | [default to null]
**RemitToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [default to null]
**ShipToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**ShipFromParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**BillToParty** | [***PartyIdentification**](PartyIdentification.md) |  | [optional] [default to null]
**PaymentTerms** | [***PaymentTerms**](PaymentTerms.md) |  | [optional] [default to null]
**InvoiceTotal** | [***Money**](Money.md) |  | [default to null]
**TaxDetails** | [**[]TaxDetails**](TaxDetails.md) | Total tax amount details for all line items. | [optional] [default to null]
**AdditionalDetails** | [**[]AdditionalDetails**](AdditionalDetails.md) | Additional details provided by the selling party, for tax related or other purposes. | [optional] [default to null]
**ChargeDetails** | [**[]ChargeDetails**](ChargeDetails.md) | Total charge amount details for all line items. | [optional] [default to null]
**AllowanceDetails** | [**[]AllowanceDetails**](AllowanceDetails.md) | Total allowance amount details for all line items. | [optional] [default to null]
**Items** | [**[]InvoiceItem**](InvoiceItem.md) | The list of invoice items. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

