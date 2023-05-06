# {{classname}}

All URIs are relative to *https://sandbox.sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateOrderScenarios**](VendorDFSandboxApi.md#GenerateOrderScenarios) | **Post** /vendor/directFulfillment/sandbox/2021-10-28/orders | 

# **GenerateOrderScenarios**
> TransactionReference GenerateOrderScenarios(ctx, body)


Submits a request to generate test order data for Vendor Direct Fulfillment API entities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GenerateOrderScenarioRequest**](GenerateOrderScenarioRequest.md)|  | 

### Return type

[**TransactionReference**](TransactionReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

