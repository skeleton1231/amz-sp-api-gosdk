# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFeed**](FeedsApi.md#CancelFeed) | **Delete** /feeds/2020-09-04/feeds/{feedId} | 
[**CreateFeed**](FeedsApi.md#CreateFeed) | **Post** /feeds/2020-09-04/feeds | 
[**CreateFeedDocument**](FeedsApi.md#CreateFeedDocument) | **Post** /feeds/2020-09-04/documents | 
[**GetFeed**](FeedsApi.md#GetFeed) | **Get** /feeds/2020-09-04/feeds/{feedId} | 
[**GetFeedDocument**](FeedsApi.md#GetFeedDocument) | **Get** /feeds/2020-09-04/documents/{feedDocumentId} | 
[**GetFeeds**](FeedsApi.md#GetFeeds) | **Get** /feeds/2020-09-04/feeds | 

# **CancelFeed**
> CancelFeedResponse CancelFeed(ctx, feedId)


Effective June 27, 2023, the `cancelFeed` operation will no longer be available in the Selling Partner API for Feeds v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Feeds v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feedId** | **string**| The identifier for the feed. This identifier is unique only in combination with a seller ID. | 

### Return type

[**CancelFeedResponse**](CancelFeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFeed**
> CreateFeedResponse CreateFeed(ctx, body)


Effective June 27, 2023, the `createFeed` operation will no longer be available in the Selling Partner API for Feeds v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Feeds v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference) to avoid service disruption.

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


Effective June 27, 2023, the `createFeedDocument` operation will no longer be available in the Selling Partner API for Feeds v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Feeds v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference) to avoid service disruption.

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
> GetFeedResponse GetFeed(ctx, feedId)


Effective June 27, 2023, the `getFeed` operation will no longer be available in the Selling Partner API for Feeds v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Feeds v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feedId** | **string**| The identifier for the feed. This identifier is unique only in combination with a seller ID. | 

### Return type

[**GetFeedResponse**](GetFeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeedDocument**
> GetFeedDocumentResponse GetFeedDocument(ctx, feedDocumentId)


Effective June 27, 2023, the `getFeedDocument` operation will no longer be available in the Selling Partner API for Feeds v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Feeds v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **feedDocumentId** | **string**| The identifier of the feed document. | 

### Return type

[**GetFeedDocumentResponse**](GetFeedDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeeds**
> GetFeedsResponse GetFeeds(ctx, optional)


Effective June 27, 2023, the `getFeeds` operation will no longer be available in the Selling Partner API for Feeds v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Feeds v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference) to avoid service disruption.

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

