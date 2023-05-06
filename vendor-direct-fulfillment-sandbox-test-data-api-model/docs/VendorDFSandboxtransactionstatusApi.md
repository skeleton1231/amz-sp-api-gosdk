# {{classname}}

All URIs are relative to *https://sandbox.sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrderScenarios**](VendorDFSandboxtransactionstatusApi.md#GetOrderScenarios) | **Get** /vendor/directFulfillment/sandbox/2021-10-28/transactions/{transactionId} | 

# **GetOrderScenarios**
> TransactionStatus GetOrderScenarios(ctx, transactionId)


Returns the status of the transaction indicated by the specified transactionId. If the transaction was successful, also returns the requested test order data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionId** | **string**| The transaction identifier returned in the response to the generateOrderScenarios operation. | 

### Return type

[**TransactionStatus**](TransactionStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

