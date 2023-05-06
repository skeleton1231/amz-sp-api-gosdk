# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrder**](VendorOrdersApi.md#GetOrder) | **Get** /vendor/directFulfillment/orders/v1/purchaseOrders/{purchaseOrderNumber} | 
[**GetOrders**](VendorOrdersApi.md#GetOrders) | **Get** /vendor/directFulfillment/orders/v1/purchaseOrders | 
[**SubmitAcknowledgement**](VendorOrdersApi.md#SubmitAcknowledgement) | **Post** /vendor/directFulfillment/orders/v1/acknowledgements | 

# **GetOrder**
> GetOrderResponse GetOrder(ctx, purchaseOrderNumber)


Returns purchase order information for the purchaseOrderNumber that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purchaseOrderNumber** | **string**| The order identifier for the purchase order that you want. Formatting Notes: alpha-numeric code. | 

### Return type

[**GetOrderResponse**](GetOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrders**
> GetOrdersResponse GetOrders(ctx, createdAfter, createdBefore, optional)


Returns a list of purchase orders created during the time frame that you specify. You define the time frame using the createdAfter and createdBefore parameters. You must use both parameters. You can choose to get only the purchase order numbers by setting the includeDetails parameter to false. In that case, the operation returns a list of purchase order numbers. You can then call the getOrder operation to return the details of a specific order.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createdAfter** | **time.Time**| Purchase orders that became available after this date and time will be included in the result. Must be in ISO-8601 date/time format. | 
  **createdBefore** | **time.Time**| Purchase orders that became available before this date and time will be included in the result. Must be in ISO-8601 date/time format. | 
 **optional** | ***VendorOrdersApiGetOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorOrdersApiGetOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipFromPartyId** | **optional.String**| The vendor warehouse identifier for the fulfillment warehouse. If not specified, the result will contain orders for all warehouses. | 
 **status** | **optional.String**| Returns only the purchase orders that match the specified status. If not specified, the result will contain orders that match any status. | 
 **limit** | **optional.Int64**| The limit to the number of purchase orders returned. | 
 **sortOrder** | **optional.String**| Sort the list in ascending or descending order by order creation date. | 
 **nextToken** | **optional.String**| Used for pagination when there are more orders than the specified result size limit. The token value is returned in the previous API call. | 
 **includeDetails** | **optional.String**| When true, returns the complete purchase order details. Otherwise, only purchase order numbers are returned. | [default to true]

### Return type

[**GetOrdersResponse**](GetOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, payload

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitAcknowledgement**
> SubmitAcknowledgementResponse SubmitAcknowledgement(ctx, body)


Submits acknowledgements for one or more purchase orders.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitAcknowledgementRequest**](SubmitAcknowledgementRequest.md)|  | 

### Return type

[**SubmitAcknowledgementResponse**](SubmitAcknowledgementResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

