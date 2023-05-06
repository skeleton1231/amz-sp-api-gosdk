# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFinancialEventGroups**](DefaultApi.md#ListFinancialEventGroups) | **Get** /finances/v0/financialEventGroups | 
[**ListFinancialEvents**](DefaultApi.md#ListFinancialEvents) | **Get** /finances/v0/financialEvents | 
[**ListFinancialEventsByGroupId**](DefaultApi.md#ListFinancialEventsByGroupId) | **Get** /finances/v0/financialEventGroups/{eventGroupId}/financialEvents | 
[**ListFinancialEventsByOrderId**](DefaultApi.md#ListFinancialEventsByOrderId) | **Get** /finances/v0/orders/{orderId}/financialEvents | 

# **ListFinancialEventGroups**
> ListFinancialEventGroupsResponse ListFinancialEventGroups(ctx, optional)


Returns financial event groups for a given date range.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiListFinancialEventGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListFinancialEventGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxResultsPerPage** | **optional.Int32**| The maximum number of results to return per page. If the response exceeds the maximum number of transactions or 10 MB, the API responds with &#x27;InvalidInput&#x27;. | [default to 100]
 **financialEventGroupStartedBefore** | **optional.Time**| A date used for selecting financial event groups that opened before (but not at) a specified date and time, in ISO 8601 format. The date-time  must be later than FinancialEventGroupStartedAfter and no later than two minutes before the request was submitted. If FinancialEventGroupStartedAfter and FinancialEventGroupStartedBefore are more than 180 days apart, no financial event groups are returned. | 
 **financialEventGroupStartedAfter** | **optional.Time**| A date used for selecting financial event groups that opened after (or at) a specified date and time, in ISO 8601 format. The date-time must be no later than two minutes before the request was submitted. | 
 **nextToken** | **optional.String**| A string token returned in the response of your previous request. | 

### Return type

[**ListFinancialEventGroupsResponse**](ListFinancialEventGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFinancialEvents**
> ListFinancialEventsResponse ListFinancialEvents(ctx, optional)


Returns financial events for the specified data range.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiListFinancialEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListFinancialEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxResultsPerPage** | **optional.Int32**| The maximum number of results to return per page. If the response exceeds the maximum number of transactions or 10 MB, the API responds with &#x27;InvalidInput&#x27;. | [default to 100]
 **postedAfter** | **optional.Time**| A date used for selecting financial events posted after (or at) a specified time. The date-time must be no later than two minutes before the request was submitted, in ISO 8601 date time format. | 
 **postedBefore** | **optional.Time**| A date used for selecting financial events posted before (but not at) a specified time. The date-time must be later than PostedAfter and no later than two minutes before the request was submitted, in ISO 8601 date time format. If PostedAfter and PostedBefore are more than 180 days apart, no financial events are returned. You must specify the PostedAfter parameter if you specify the PostedBefore parameter. Default: Now minus two minutes. | 
 **nextToken** | **optional.String**| A string token returned in the response of your previous request. | 

### Return type

[**ListFinancialEventsResponse**](ListFinancialEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFinancialEventsByGroupId**
> ListFinancialEventsResponse ListFinancialEventsByGroupId(ctx, eventGroupId, optional)


Returns all financial events for the specified financial event group.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventGroupId** | **string**| The identifier of the financial event group to which the events belong. | 
 **optional** | ***DefaultApiListFinancialEventsByGroupIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListFinancialEventsByGroupIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResultsPerPage** | **optional.Int32**| The maximum number of results to return per page. If the response exceeds the maximum number of transactions or 10 MB, the API responds with &#x27;InvalidInput&#x27;. | [default to 100]
 **postedAfter** | **optional.Time**| A date used for selecting financial events posted after (or at) a specified time. The date-time **must** be more than two minutes before the time of the request, in ISO 8601 date time format. | 
 **postedBefore** | **optional.Time**| A date used for selecting financial events posted before (but not at) a specified time. The date-time must be later than &#x60;PostedAfter&#x60; and no later than two minutes before the request was submitted, in ISO 8601 date time format. If &#x60;PostedAfter&#x60; and &#x60;PostedBefore&#x60; are more than 180 days apart, no financial events are returned. You must specify the &#x60;PostedAfter&#x60; parameter if you specify the &#x60;PostedBefore&#x60; parameter. Default: Now minus two minutes. | 
 **nextToken** | **optional.String**| A string token returned in the response of your previous request. | 

### Return type

[**ListFinancialEventsResponse**](ListFinancialEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFinancialEventsByOrderId**
> ListFinancialEventsResponse ListFinancialEventsByOrderId(ctx, orderId, optional)


Returns all financial events for the specified order.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| An Amazon-defined order identifier, in 3-7-7 format. | 
 **optional** | ***DefaultApiListFinancialEventsByOrderIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListFinancialEventsByOrderIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResultsPerPage** | **optional.Int32**| The maximum number of results to return per page. If the response exceeds the maximum number of transactions or 10 MB, the API responds with &#x27;InvalidInput&#x27;. | [default to 100]
 **nextToken** | **optional.String**| A string token returned in the response of your previous request. | 

### Return type

[**ListFinancialEventsResponse**](ListFinancialEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

