# ServiceFeeEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonOrderId** | **string** | An Amazon-defined identifier for an order. | [optional] [default to null]
**FeeReason** | **string** | A short description of the service fee reason. | [optional] [default to null]
**FeeList** | [***[]FeeComponent**](array.md) |  | [optional] [default to null]
**SellerSKU** | **string** | The seller SKU of the item. The seller SKU is qualified by the seller&#x27;s seller ID, which is included with every call to the Selling Partner API. | [optional] [default to null]
**FnSKU** | **string** | A unique identifier assigned by Amazon to products stored in and fulfilled from an Amazon fulfillment center. | [optional] [default to null]
**FeeDescription** | **string** | A short description of the service fee event. | [optional] [default to null]
**ASIN** | **string** | The Amazon Standard Identification Number (ASIN) of the item. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

