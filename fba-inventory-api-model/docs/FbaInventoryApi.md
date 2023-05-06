# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInventorySummaries**](FbaInventoryApi.md#GetInventorySummaries) | **Get** /fba/inventory/v1/summaries | 

# **GetInventorySummaries**
> GetInventorySummariesResponse GetInventorySummaries(ctx, granularityType, granularityId, marketplaceIds, optional)


Returns a list of inventory summaries. The summaries returned depend on the presence or absence of the `startDateTime`, `sellerSkus` and `sellerSku` parameters:  - All inventory summaries with available details are returned when the `startDateTime`, `sellerSkus` and `sellerSku` parameters are omitted. - When `startDateTime` is provided, the operation returns inventory summaries that have had changes after the date and time specified. The `sellerSkus` and `sellerSku` parameters are ignored. **Important:** To avoid errors, use both `startDateTime` and `nextToken` to get the next page of inventory summaries that have changed after the date and time specified. - When the `sellerSkus` parameter is provided, the operation returns inventory summaries for only the specified `sellerSkus`. The `sellerSku` parameter is ignored. - When the `sellerSku` parameter is provided, the operation returns inventory summaries for only the specified `sellerSku`.  **Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 2 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **granularityType** | **string**| The granularity type for the inventory aggregation level. | 
  **granularityId** | **string**| The granularity ID for the inventory aggregation level. | 
  **marketplaceIds** | [**[]string**](string.md)| The marketplace ID for the marketplace for which to return inventory summaries. | 
 **optional** | ***FbaInventoryApiGetInventorySummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FbaInventoryApiGetInventorySummariesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **details** | **optional.Bool**| true to return inventory summaries with additional summarized inventory details and quantities. Otherwise, returns inventory summaries only (default value). | [default to false]
 **startDateTime** | **optional.Time**| A start date and time in ISO8601 format. If specified, all inventory summaries that have changed since then are returned. You must specify a date and time that is no earlier than 18 months prior to the date and time when you call the API. Note: Changes in inboundWorkingQuantity, inboundShippedQuantity and inboundReceivingQuantity are not detected. | 
 **sellerSkus** | [**optional.Interface of []string**](string.md)| A list of seller SKUs for which to return inventory summaries. You may specify up to 50 SKUs. | 
 **sellerSku** | **optional.String**| A single seller SKU used for querying the specified seller SKU inventory summaries. | 
 **nextToken** | **optional.String**| String token returned in the response of your previous request. The string token will expire 30 seconds after being created. | 

### Return type

[**GetInventorySummariesResponse**](GetInventorySummariesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

