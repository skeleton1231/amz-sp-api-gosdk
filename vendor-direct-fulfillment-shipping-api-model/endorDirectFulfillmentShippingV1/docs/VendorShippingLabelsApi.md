# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShippingLabel**](VendorShippingLabelsApi.md#GetShippingLabel) | **Get** /vendor/directFulfillment/shipping/v1/shippingLabels/{purchaseOrderNumber} | 
[**GetShippingLabels**](VendorShippingLabelsApi.md#GetShippingLabels) | **Get** /vendor/directFulfillment/shipping/v1/shippingLabels | 
[**SubmitShippingLabelRequest**](VendorShippingLabelsApi.md#SubmitShippingLabelRequest) | **Post** /vendor/directFulfillment/shipping/v1/shippingLabels | 

# **GetShippingLabel**
> GetShippingLabelResponse GetShippingLabel(ctx, purchaseOrderNumber)


Returns a shipping label for the purchaseOrderNumber that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purchaseOrderNumber** | **string**| The purchase order number for which you want to return the shipping label. It should be the same purchaseOrderNumber as received in the order. | 

### Return type

[**GetShippingLabelResponse**](GetShippingLabelResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippingLabels**
> GetShippingLabelListResponse GetShippingLabels(ctx, createdAfter, createdBefore, optional)


Returns a list of shipping labels created during the time frame that you specify. You define that time frame using the createdAfter and createdBefore parameters. You must use both of these parameters. The date range to search must not be more than 7 days.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createdAfter** | **time.Time**| Shipping labels that became available after this date and time will be included in the result. Must be in ISO-8601 date/time format. | 
  **createdBefore** | **time.Time**| Shipping labels that became available before this date and time will be included in the result. Must be in ISO-8601 date/time format. | 
 **optional** | ***VendorShippingLabelsApiGetShippingLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorShippingLabelsApiGetShippingLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipFromPartyId** | **optional.String**| The vendor warehouseId for order fulfillment. If not specified, the result will contain orders for all warehouses. | 
 **limit** | **optional.Int32**| The limit to the number of records returned. | 
 **sortOrder** | **optional.String**| Sort ASC or DESC by order creation date. | [default to ASC]
 **nextToken** | **optional.String**| Used for pagination when there are more ship labels than the specified result size limit. The token value is returned in the previous API call. | 

### Return type

[**GetShippingLabelListResponse**](GetShippingLabelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, payload

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitShippingLabelRequest**
> SubmitShippingLabelsResponse SubmitShippingLabelRequest(ctx, body)


Creates a shipping label for a purchase order and returns a transactionId for reference.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitShippingLabelsRequest**](SubmitShippingLabelsRequest.md)|  | 

### Return type

[**SubmitShippingLabelsResponse**](SubmitShippingLabelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

