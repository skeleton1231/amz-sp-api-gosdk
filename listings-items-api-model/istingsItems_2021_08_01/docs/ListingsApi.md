# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteListingsItem**](ListingsApi.md#DeleteListingsItem) | **Delete** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**GetListingsItem**](ListingsApi.md#GetListingsItem) | **Get** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**PatchListingsItem**](ListingsApi.md#PatchListingsItem) | **Patch** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**PutListingsItem**](ListingsApi.md#PutListingsItem) | **Put** /listings/2021-08-01/items/{sellerId}/{sku} | 

# **DeleteListingsItem**
> ListingsItemSubmissionResponse DeleteListingsItem(ctx, sellerId, sku, marketplaceIds, optional)


Delete a listings item for a selling partner.  **Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sellerId** | **string**| A selling partner identifier, such as a merchant account or vendor code. | 
  **sku** | **string**| A selling partner provided identifier for an Amazon listing. | 
  **marketplaceIds** | [**[]string**](string.md)| A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **optional** | ***ListingsApiDeleteListingsItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListingsApiDeleteListingsItemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issueLocale** | **optional.String**| A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. Localized messages default to \&quot;en_US\&quot; when a localization is not available in the specified locale. | 

### Return type

[**ListingsItemSubmissionResponse**](ListingsItemSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListingsItem**
> Item GetListingsItem(ctx, sellerId, sku, marketplaceIds, optional)


Returns details about a listings item for a selling partner.  **Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sellerId** | **string**| A selling partner identifier, such as a merchant account or vendor code. | 
  **sku** | **string**| A selling partner provided identifier for an Amazon listing. | 
  **marketplaceIds** | [**[]string**](string.md)| A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **optional** | ***ListingsApiGetListingsItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListingsApiGetListingsItemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issueLocale** | **optional.String**| A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. Localized messages default to \&quot;en_US\&quot; when a localization is not available in the specified locale. | 
 **includedData** | [**optional.Interface of []string**](string.md)| A comma-delimited list of data sets to include in the response. Default: summaries. | 

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchListingsItem**
> ListingsItemSubmissionResponse PatchListingsItem(ctx, body, sellerId, sku, marketplaceIds, optional)


Partially update (patch) a listings item for a selling partner. Only top-level listings item attributes can be patched. Patching nested attributes is not supported.  **Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListingsItemPatchRequest**](ListingsItemPatchRequest.md)| The request body schema for the patchListingsItem operation. | 
  **sellerId** | **string**| A selling partner identifier, such as a merchant account or vendor code. | 
  **sku** | **string**| A selling partner provided identifier for an Amazon listing. | 
  **marketplaceIds** | [**[]string**](string.md)| A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **optional** | ***ListingsApiPatchListingsItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListingsApiPatchListingsItemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **issueLocale** | **optional.**| A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. Localized messages default to \&quot;en_US\&quot; when a localization is not available in the specified locale. | 

### Return type

[**ListingsItemSubmissionResponse**](ListingsItemSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutListingsItem**
> ListingsItemSubmissionResponse PutListingsItem(ctx, body, sellerId, sku, marketplaceIds, optional)


Creates a new or fully-updates an existing listings item for a selling partner.  **Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListingsItemPutRequest**](ListingsItemPutRequest.md)| The request body schema for the putListingsItem operation. | 
  **sellerId** | **string**| A selling partner identifier, such as a merchant account or vendor code. | 
  **sku** | **string**| A selling partner provided identifier for an Amazon listing. | 
  **marketplaceIds** | [**[]string**](string.md)| A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **optional** | ***ListingsApiPutListingsItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListingsApiPutListingsItemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **issueLocale** | **optional.**| A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. Localized messages default to \&quot;en_US\&quot; when a localization is not available in the specified locale. | 

### Return type

[**ListingsItemSubmissionResponse**](ListingsItemSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

