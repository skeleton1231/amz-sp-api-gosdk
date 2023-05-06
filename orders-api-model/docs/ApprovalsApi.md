# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrderItemsApprovals**](ApprovalsApi.md#GetOrderItemsApprovals) | **Get** /orders/v0/orders/{orderId}/approvals | 
[**UpdateOrderItemsApprovals**](ApprovalsApi.md#UpdateOrderItemsApprovals) | **Post** /orders/v0/orders/{orderId}/approvals | 

# **GetOrderItemsApprovals**
> GetOrderApprovalsResponse GetOrderItemsApprovals(ctx, orderId, optional)


Returns detailed order items approvals information for the order specified. If NextToken is provided, it's used to retrieve the next page of order items approvals.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 0.5 | 30 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| An Amazon-defined order identifier, in 3-7-7 format, e.g. 933-1671587-0818628. | 
 **optional** | ***ApprovalsApiGetOrderItemsApprovalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApprovalsApiGetOrderItemsApprovalsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextToken** | **optional.String**| A string token returned in the response of your previous request. | 
 **itemApprovalTypes** | [**optional.Interface of []ItemApprovalType**](ItemApprovalType.md)| When set, only return approvals for items which approval type is contained in the specified approval types. | 
 **itemApprovalStatus** | [**optional.Interface of []ItemApprovalStatus**](ItemApprovalStatus.md)| When set, only return approvals that contain items which approval status is contained in the specified approval status. | 

### Return type

[**GetOrderApprovalsResponse**](GetOrderApprovalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrderItemsApprovals**
> UpdateOrderItemsApprovals(ctx, body, orderId)


Update the order items approvals for an order that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 15 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOrderApprovalsRequest**](UpdateOrderApprovalsRequest.md)| The request body for the updateOrderItemsApprovals operation. | 
  **orderId** | **string**| An Amazon-defined order identifier, in 3-7-7 format. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

