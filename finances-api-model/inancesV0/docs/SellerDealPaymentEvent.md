# SellerDealPaymentEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**DealId** | **string** | The unique identifier of the deal. | [optional] [default to null]
**DealDescription** | **string** | The internal description of the deal. | [optional] [default to null]
**EventType** | **string** | The type of event: SellerDealComplete. | [optional] [default to null]
**FeeType** | **string** | The type of fee: RunLightningDealFee. | [optional] [default to null]
**FeeAmount** | [***Currency**](Currency.md) |  | [optional] [default to null]
**TaxAmount** | [***Currency**](Currency.md) |  | [optional] [default to null]
**TotalAmount** | [***Currency**](Currency.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

