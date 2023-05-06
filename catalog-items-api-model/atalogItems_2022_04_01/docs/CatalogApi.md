# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogItem**](CatalogApi.md#GetCatalogItem) | **Get** /catalog/2022-04-01/items/{asin} | 
[**SearchCatalogItems**](CatalogApi.md#SearchCatalogItems) | **Get** /catalog/2022-04-01/items | 

# **GetCatalogItem**
> Item GetCatalogItem(ctx, asin, marketplaceIds, optional)


Retrieves details for an item in the Amazon catalog.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 2 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to the [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asin** | **string**| The Amazon Standard Identification Number (ASIN) of the item. | 
  **marketplaceIds** | [**[]string**](string.md)| A comma-delimited list of Amazon marketplace identifiers. Data sets in the response contain data only for the specified marketplaces. | 
 **optional** | ***CatalogApiGetCatalogItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiGetCatalogItemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includedData** | [**optional.Interface of []string**](string.md)| A comma-delimited list of data sets to include in the response. Default: &#x60;summaries&#x60;. | 
 **locale** | **optional.String**| Locale for retrieving localized summaries. Defaults to the primary locale of the marketplace. | 

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCatalogItems**
> ItemSearchResults SearchCatalogItems(ctx, marketplaceIds, optional)


Search for and return a list of Amazon catalog items and associated information either by identifier or by keywords.  **Usage Plans:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 2 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to the [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceIds** | [**[]string**](string.md)| A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **optional** | ***CatalogApiSearchCatalogItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiSearchCatalogItemsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identifiers** | [**optional.Interface of []string**](string.md)| A comma-delimited list of product identifiers to search the Amazon catalog for. **Note:** Cannot be used with &#x60;keywords&#x60;. | 
 **identifiersType** | **optional.String**| Type of product identifiers to search the Amazon catalog for. **Note:** Required when &#x60;identifiers&#x60; are provided. | 
 **includedData** | [**optional.Interface of []string**](string.md)| A comma-delimited list of data sets to include in the response. Default: &#x60;summaries&#x60;. | 
 **locale** | **optional.String**| Locale for retrieving localized summaries. Defaults to the primary locale of the marketplace. | 
 **sellerId** | **optional.String**| A selling partner identifier, such as a seller account or vendor code. **Note:** Required when &#x60;identifiersType&#x60; is &#x60;SKU&#x60;. | 
 **keywords** | [**optional.Interface of []string**](string.md)| A comma-delimited list of words to search the Amazon catalog for. **Note:** Cannot be used with &#x60;identifiers&#x60;. | 
 **brandNames** | [**optional.Interface of []string**](string.md)| A comma-delimited list of brand names to limit the search for &#x60;keywords&#x60;-based queries. **Note:** Cannot be used with &#x60;identifiers&#x60;. | 
 **classificationIds** | [**optional.Interface of []string**](string.md)| A comma-delimited list of classification identifiers to limit the search for &#x60;keywords&#x60;-based queries. **Note:** Cannot be used with &#x60;identifiers&#x60;. | 
 **pageSize** | **optional.Int32**| Number of results to be returned per page. | [default to 10]
 **pageToken** | **optional.String**| A token to fetch a certain page when there are multiple pages worth of results. | 
 **keywordsLocale** | **optional.String**| The language of the keywords provided for &#x60;keywords&#x60;-based queries. Defaults to the primary locale of the marketplace. **Note:** Cannot be used with &#x60;identifiers&#x60;. | 

### Return type

[**ItemSearchResults**](ItemSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

