# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListingsRestrictions**](ListingsApi.md#GetListingsRestrictions) | **Get** /listings/2021-08-01/restrictions | 

# **GetListingsRestrictions**
> RestrictionList GetListingsRestrictions(ctx, asin, sellerId, marketplaceIds, optional)


Returns listing restrictions for an item in the Amazon Catalog.   **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asin** | **string**| The Amazon Standard Identification Number (ASIN) of the item. | 
  **sellerId** | **string**| A selling partner identifier, such as a merchant account. | 
  **marketplaceIds** | [**[]string**](string.md)| A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **optional** | ***ListingsApiGetListingsRestrictionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListingsApiGetListingsRestrictionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **conditionType** | **optional.String**| The condition used to filter restrictions. | 
 **reasonLocale** | **optional.String**| A locale for reason text localization. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. Localized messages default to \&quot;en_US\&quot; when a localization is not available in the specified locale. | 

### Return type

[**RestrictionList**](RestrictionList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

