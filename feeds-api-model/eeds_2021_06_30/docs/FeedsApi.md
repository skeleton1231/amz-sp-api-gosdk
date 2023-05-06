# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFeed**](FeedsApi.md#CancelFeed) | **Delete** /feeds/2021-06-30/feeds/{feedId} | 
[**CreateFeed**](FeedsApi.md#CreateFeed) | **Post** /feeds/2021-06-30/feeds | 
[**CreateFeedDocument**](FeedsApi.md#CreateFeedDocument) | **Post** /feeds/2021-06-30/documents | 
[**GetFeed**](FeedsApi.md#GetFeed) | **Get** /feeds/2021-06-30/feeds/{feedId} | 
[**GetFeedDocument**](FeedsApi.md#GetFeedDocument) | **Get** /feeds/2021-06-30/documents/{feedDocumentId} | 
[**GetFeeds**](FeedsApi.md#GetFeeds) | **Get** /feeds/2021-06-30/feeds | 

# **CancelFeed**
> CancelFeed(ctx, feedId)


Cancels the feed that you specify. Only feeds with processingStatus=IN_QUEUE can be cancelled. Cancelled feeds are returned in subsequent calls to the getFeed and getFeeds operations.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 15 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feedId** | **string**| The identifier for the feed. This identifier is unique only in combination with a seller ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFeed**
> CreateFeedResponse CreateFeed(ctx, body)


Creates a feed. Upload the contents of the feed document before calling this operation.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.0083 | 15 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFeedSpecification**](CreateFeedSpecification.md)|  | 

### Return type

[**CreateFeedResponse**](CreateFeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFeedDocument**
> CreateFeedDocumentResponse CreateFeedDocument(ctx, body)


Creates a feed document for the feed type that you specify. This operation returns a presigned URL for uploading the feed document contents. It also returns a feedDocumentId value that you can pass in with a subsequent call to the createFeed operation.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 15 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFeedDocumentSpecification**](CreateFeedDocumentSpecification.md)|  | 

### Return type

[**CreateFeedDocumentResponse**](CreateFeedDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeed**
> Feed GetFeed(ctx, feedId)


Returns feed details (including the resultDocumentId, if available) for the feed that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 15 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feedId** | **string**| The identifier for the feed. This identifier is unique only in combination with a seller ID. | 

### Return type

[**Feed**](Feed.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeedDocument**
> FeedDocument GetFeedDocument(ctx, feedDocumentId)


Returns the information required for retrieving a feed document's contents.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.0222 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feedDocumentId** | **string**| The identifier of the feed document. | 

### Return type

[**FeedDocument**](FeedDocument.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeeds**
> GetFeedsResponse GetFeeds(ctx, optional)


Returns feed details for the feeds that match the filters that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.0222 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FeedsApiGetFeedsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedsApiGetFeedsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedTypes** | [**optional.Interface of []string**](string.md)| A list of feed types used to filter feeds. When feedTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either feedTypes or nextToken is required. | 
 **marketplaceIds** | [**optional.Interface of []string**](string.md)| A list of marketplace identifiers used to filter feeds. The feeds returned will match at least one of the marketplaces that you specify. | 
 **pageSize** | **optional.Int32**| The maximum number of feeds to return in a single call. | [default to 10]
 **processingStatuses** | [**optional.Interface of []string**](string.md)| A list of processing statuses used to filter feeds. | 
 **createdSince** | **optional.Time**| The earliest feed creation date and time for feeds included in the response, in ISO 8601 format. The default is 90 days ago. Feeds are retained for a maximum of 90 days. | 
 **createdUntil** | **optional.Time**| The latest feed creation date and time for feeds included in the response, in ISO 8601 format. The default is now. | 
 **nextToken** | **optional.String**| A string token returned in the response to your previous request. nextToken is returned when the number of results exceeds the specified pageSize value. To get the next page of results, call the getFeeds operation and include this token as the only parameter. Specifying nextToken with any other parameters will cause the request to fail. | 

### Return type

[**GetFeedsResponse**](GetFeedsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

