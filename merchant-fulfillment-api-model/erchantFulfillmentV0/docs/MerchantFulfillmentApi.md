# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelShipment**](MerchantFulfillmentApi.md#CancelShipment) | **Delete** /mfn/v0/shipments/{shipmentId} | 
[**CancelShipmentOld**](MerchantFulfillmentApi.md#CancelShipmentOld) | **Put** /mfn/v0/shipments/{shipmentId}/cancel | 
[**CreateShipment**](MerchantFulfillmentApi.md#CreateShipment) | **Post** /mfn/v0/shipments | 
[**GetAdditionalSellerInputs**](MerchantFulfillmentApi.md#GetAdditionalSellerInputs) | **Post** /mfn/v0/additionalSellerInputs | 
[**GetAdditionalSellerInputsOld**](MerchantFulfillmentApi.md#GetAdditionalSellerInputsOld) | **Post** /mfn/v0/sellerInputs | 
[**GetEligibleShipmentServices**](MerchantFulfillmentApi.md#GetEligibleShipmentServices) | **Post** /mfn/v0/eligibleShippingServices | 
[**GetEligibleShipmentServicesOld**](MerchantFulfillmentApi.md#GetEligibleShipmentServicesOld) | **Post** /mfn/v0/eligibleServices | 
[**GetShipment**](MerchantFulfillmentApi.md#GetShipment) | **Get** /mfn/v0/shipments/{shipmentId} | 

# **CancelShipment**
> CancelShipmentResponse CancelShipment(ctx, shipmentId)


Cancel the shipment indicated by the specified shipment identifier.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| The Amazon-defined shipment identifier for the shipment to cancel. | 

### Return type

[**CancelShipmentResponse**](CancelShipmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelShipmentOld**
> CancelShipmentResponse CancelShipmentOld(ctx, shipmentId)


Cancel the shipment indicated by the specified shipment identifer.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| The Amazon-defined shipment identifier for the shipment to cancel. | 

### Return type

[**CancelShipmentResponse**](CancelShipmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateShipment**
> CreateShipmentResponse CreateShipment(ctx, body)


Create a shipment with the information provided.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateShipmentRequest**](CreateShipmentRequest.md)|  | 

### Return type

[**CreateShipmentResponse**](CreateShipmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdditionalSellerInputs**
> GetAdditionalSellerInputsResponse GetAdditionalSellerInputs(ctx, body)


Gets a list of additional seller inputs required for a ship method. This is generally used for international shipping.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetAdditionalSellerInputsRequest**](GetAdditionalSellerInputsRequest.md)|  | 

### Return type

[**GetAdditionalSellerInputsResponse**](GetAdditionalSellerInputsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdditionalSellerInputsOld**
> GetAdditionalSellerInputsResponse GetAdditionalSellerInputsOld(ctx, body)


Get a list of additional seller inputs required for a ship method. This is generally used for international shipping.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetAdditionalSellerInputsRequest**](GetAdditionalSellerInputsRequest.md)|  | 

### Return type

[**GetAdditionalSellerInputsResponse**](GetAdditionalSellerInputsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEligibleShipmentServices**
> GetEligibleShipmentServicesResponse GetEligibleShipmentServices(ctx, body)


Returns a list of shipping service offers that satisfy the specified shipment request details.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetEligibleShipmentServicesRequest**](GetEligibleShipmentServicesRequest.md)|  | 

### Return type

[**GetEligibleShipmentServicesResponse**](GetEligibleShipmentServicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEligibleShipmentServicesOld**
> GetEligibleShipmentServicesResponse GetEligibleShipmentServicesOld(ctx, body)


Returns a list of shipping service offers that satisfy the specified shipment request details.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetEligibleShipmentServicesRequest**](GetEligibleShipmentServicesRequest.md)|  | 

### Return type

[**GetEligibleShipmentServicesResponse**](GetEligibleShipmentServicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShipment**
> GetShipmentResponse GetShipment(ctx, shipmentId)


Returns the shipment information for an existing shipment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| The Amazon-defined shipment identifier for the shipment. | 

### Return type

[**GetShipmentResponse**](GetShipmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

