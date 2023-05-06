# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogItem**](CatalogApi.md#GetCatalogItem) | **Get** /catalog/v0/items/{asin} | 
[**ListCatalogCategories**](CatalogApi.md#ListCatalogCategories) | **Get** /catalog/v0/categories | 
[**ListCatalogItems**](CatalogApi.md#ListCatalogItems) | **Get** /catalog/v0/items | 

# **GetCatalogItem**
> GetCatalogItemResponse GetCatalogItem(ctx, marketplaceId, asin)


Effective September 30, 2022, the `getCatalogItem` operation will no longer be available in the Selling Partner API for Catalog Items v0. This operation is available in the latest version of the [Selling Partner API for Catalog Items v2022-04-01](doc:catalog-items-api-v2022-04-01-reference). Integrations that rely on this operation should migrate to the latest version to avoid service disruption.  _Note:_ The [`listCatalogCategories`](#get-catalogv0categories) operation is not being deprecated and you can continue to make calls to it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace for the item. | 
  **asin** | **string**| The Amazon Standard Identification Number (ASIN) of the item. | 

### Return type

[**GetCatalogItemResponse**](GetCatalogItemResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCatalogCategories**
> ListCatalogCategoriesResponse ListCatalogCategories(ctx, marketplaceId, optional)


Returns the parent categories to which an item belongs, based on the specified ASIN or SellerSKU.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 2 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace for the item. | 
 **optional** | ***CatalogApiListCatalogCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiListCatalogCategoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aSIN** | **optional.String**| The Amazon Standard Identification Number (ASIN) of the item. | 
 **sellerSKU** | **optional.String**| Used to identify items in the given marketplace. SellerSKU is qualified by the seller&#x27;s SellerId, which is included with every operation that you submit. | 

### Return type

[**ListCatalogCategoriesResponse**](ListCatalogCategoriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCatalogItems**
> ListCatalogItemsResponse ListCatalogItems(ctx, marketplaceId, optional)


Effective September 30, 2022, the `listCatalogItems` operation will no longer be available in the Selling Partner API for Catalog Items v0. As an alternative, `searchCatalogItems` is available in the latest version of the [Selling Partner API for Catalog Items v2022-04-01](doc:catalog-items-api-v2022-04-01-reference). Integrations that rely on the `listCatalogItems` operation should migrate to the `searchCatalogItems`operation to avoid service disruption.  _Note:_ The [`listCatalogCategories`](#get-catalogv0categories) operation is not being deprecated and you can continue to make calls to it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace for which items are returned. | 
 **optional** | ***CatalogApiListCatalogItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiListCatalogItemsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| Keyword(s) to use to search for items in the catalog. Example: &#x27;harry potter books&#x27;. | 
 **queryContextId** | **optional.String**| An identifier for the context within which the given search will be performed. A marketplace might provide mechanisms for constraining a search to a subset of potential items. For example, the retail marketplace allows queries to be constrained to a specific category. The QueryContextId parameter specifies such a subset. If it is omitted, the search will be performed using the default context for the marketplace, which will typically contain the largest set of items. | 
 **sellerSKU** | **optional.String**| Used to identify an item in the given marketplace. SellerSKU is qualified by the seller&#x27;s SellerId, which is included with every operation that you submit. | 
 **uPC** | **optional.String**| A 12-digit bar code used for retail packaging. | 
 **eAN** | **optional.String**| A European article number that uniquely identifies the catalog item, manufacturer, and its attributes. | 
 **iSBN** | **optional.String**| The unique commercial book identifier used to identify books internationally. | 
 **jAN** | **optional.String**| A Japanese article number that uniquely identifies the product, manufacturer, and its attributes. | 

### Return type

[**ListCatalogItemsResponse**](ListCatalogItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

