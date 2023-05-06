# PayWithAmazonEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerOrderId** | **string** | An order identifier that is specified by the seller. | [optional] [default to null]
**TransactionPostedDate** | [***time.Time**](time.Time.md) |  | [optional] [default to null]
**BusinessObjectType** | **string** | The type of business object. | [optional] [default to null]
**SalesChannel** | **string** | The sales channel for the transaction. | [optional] [default to null]
**Charge** | [***ChargeComponent**](ChargeComponent.md) |  | [optional] [default to null]
**FeeList** | [***[]FeeComponent**](array.md) |  | [optional] [default to null]
**PaymentAmountType** | **string** | The type of payment.  Possible values:  * Sales | [optional] [default to null]
**AmountDescription** | **string** | A short description of this payment event. | [optional] [default to null]
**FulfillmentChannel** | **string** | The fulfillment channel.  Possible values:  * AFN - Amazon Fulfillment Network (Fulfillment by Amazon)  * MFN - Merchant Fulfillment Network (self-fulfilled) | [optional] [default to null]
**StoreName** | **string** | The store name where the event occurred. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

