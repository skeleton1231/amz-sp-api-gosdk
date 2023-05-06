# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShippingLabels**](VendorShippingLabelsApi.md#CreateShippingLabels) | **Post** /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber} | 
[**GetShippingLabel**](VendorShippingLabelsApi.md#GetShippingLabel) | **Get** /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber} | 
[**GetShippingLabels**](VendorShippingLabelsApi.md#GetShippingLabels) | **Get** /vendor/directFulfillment/shipping/2021-12-28/shippingLabels | 
[**SubmitShippingLabelRequest**](VendorShippingLabelsApi.md#SubmitShippingLabelRequest) | **Post** /vendor/directFulfillment/shipping/2021-12-28/shippingLabels | 

# **CreateShippingLabels**
> ShippingLabel CreateShippingLabels(ctx, body, purchaseOrderNumber)


Creates shipping labels for a purchase order and returns the labels.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateShippingLabelsRequest**](CreateShippingLabelsRequest.md)|  | 
  **purchaseOrderNumber** | **string**| The purchase order number for which you want to return the shipping labels. It should be the same purchaseOrderNumber as received in the order. | 

### Return type

[**ShippingLabel**](ShippingLabel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippingLabel**
> ShippingLabel GetShippingLabel(ctx, purchaseOrderNumber)


Returns a shipping label for the purchaseOrderNumber that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purchaseOrderNumber** | **string**| The purchase order number for which you want to return the shipping label. It should be the same purchaseOrderNumber as received in the order. | 

### Return type

[**ShippingLabel**](ShippingLabel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShippingLabels**
> ShippingLabelList GetShippingLabels(ctx, createdAfter, createdBefore, optional)


Returns a list of shipping labels created during the time frame that you specify. You define that time frame using the createdAfter and createdBefore parameters. You must use both of these parameters. The date range to search must not be more than 7 days.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

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

[**ShippingLabelList**](ShippingLabelList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, pagination, shippingLabels

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitShippingLabelRequest**
> TransactionReference SubmitShippingLabelRequest(ctx, body)


Creates a shipping label for a purchase order and returns a transactionId for reference.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitShippingLabelsRequest**](SubmitShippingLabelsRequest.md)|  | 

### Return type

[**TransactionReference**](TransactionReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

