# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompetitivePricing**](ProductPricingApi.md#GetCompetitivePricing) | **Get** /products/pricing/v0/competitivePrice | 
[**GetItemOffers**](ProductPricingApi.md#GetItemOffers) | **Get** /products/pricing/v0/items/{Asin}/offers | 
[**GetItemOffersBatch**](ProductPricingApi.md#GetItemOffersBatch) | **Post** /batches/products/pricing/v0/itemOffers | 
[**GetListingOffers**](ProductPricingApi.md#GetListingOffers) | **Get** /products/pricing/v0/listings/{SellerSKU}/offers | 
[**GetListingOffersBatch**](ProductPricingApi.md#GetListingOffersBatch) | **Post** /batches/products/pricing/v0/listingOffers | 
[**GetPricing**](ProductPricingApi.md#GetPricing) | **Get** /products/pricing/v0/price | 

# **GetCompetitivePricing**
> GetPricingResponse GetCompetitivePricing(ctx, marketplaceId, itemType, optional)


Returns competitive pricing information for a seller's offer listings based on seller SKU or ASIN.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace for which prices are returned. | 
  **itemType** | **string**| Indicates whether ASIN values or seller SKU values are used to identify items. If you specify Asin, the information in the response will be dependent on the list of Asins you provide in the Asins parameter. If you specify Sku, the information in the response will be dependent on the list of Skus you provide in the Skus parameter. Possible values: Asin, Sku. | 
 **optional** | ***ProductPricingApiGetCompetitivePricingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductPricingApiGetCompetitivePricingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asins** | [**optional.Interface of []string**](string.md)| A list of up to twenty Amazon Standard Identification Number (ASIN) values used to identify items in the given marketplace. | 
 **skus** | [**optional.Interface of []string**](string.md)| A list of up to twenty seller SKU values used to identify items in the given marketplace. | 
 **customerType** | **optional.String**| Indicates whether to request pricing information from the point of view of Consumer or Business buyers. Default is Consumer. | 

### Return type

[**GetPricingResponse**](GetPricingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetItemOffers**
> GetOffersResponse GetItemOffers(ctx, marketplaceId, itemCondition, asin, optional)


Returns the lowest priced offers for a single item based on ASIN.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace for which prices are returned. | 
  **itemCondition** | **string**| Filters the offer listings to be considered based on item condition. Possible values: New, Used, Collectible, Refurbished, Club. | 
  **asin** | **string**| The Amazon Standard Identification Number (ASIN) of the item. | 
 **optional** | ***ProductPricingApiGetItemOffersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductPricingApiGetItemOffersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **customerType** | **optional.String**| Indicates whether to request Consumer or Business offers. Default is Consumer. | 

### Return type

[**GetOffersResponse**](GetOffersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetItemOffersBatch**
> GetItemOffersBatchResponse GetItemOffersBatch(ctx, body)


Returns the lowest priced offers for a batch of items based on ASIN.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetItemOffersBatchRequest**](GetItemOffersBatchRequest.md)|  | 

### Return type

[**GetItemOffersBatchResponse**](GetItemOffersBatchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListingOffers**
> GetOffersResponse GetListingOffers(ctx, marketplaceId, itemCondition, sellerSKU, optional)


Returns the lowest priced offers for a single SKU listing.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 2 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace for which prices are returned. | 
  **itemCondition** | **string**| Filters the offer listings based on item condition. Possible values: New, Used, Collectible, Refurbished, Club. | 
  **sellerSKU** | **string**| Identifies an item in the given marketplace. SellerSKU is qualified by the seller&#x27;s SellerId, which is included with every operation that you submit. | 
 **optional** | ***ProductPricingApiGetListingOffersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductPricingApiGetListingOffersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **customerType** | **optional.String**| Indicates whether to request Consumer or Business offers. Default is Consumer. | 

### Return type

[**GetOffersResponse**](GetOffersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListingOffersBatch**
> GetListingOffersBatchResponse GetListingOffersBatch(ctx, body)


Returns the lowest priced offers for a batch of listings by SKU.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetListingOffersBatchRequest**](GetListingOffersBatchRequest.md)|  | 

### Return type

[**GetListingOffersBatchResponse**](GetListingOffersBatchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPricing**
> GetPricingResponse GetPricing(ctx, marketplaceId, itemType, optional)


Returns pricing information for a seller's offer listings based on seller SKU or ASIN.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 0.5 | 1 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace for which prices are returned. | 
  **itemType** | **string**| Indicates whether ASIN values or seller SKU values are used to identify items. If you specify Asin, the information in the response will be dependent on the list of Asins you provide in the Asins parameter. If you specify Sku, the information in the response will be dependent on the list of Skus you provide in the Skus parameter. | 
 **optional** | ***ProductPricingApiGetPricingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductPricingApiGetPricingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asins** | [**optional.Interface of []string**](string.md)| A list of up to twenty Amazon Standard Identification Number (ASIN) values used to identify items in the given marketplace. | 
 **skus** | [**optional.Interface of []string**](string.md)| A list of up to twenty seller SKU values used to identify items in the given marketplace. | 
 **itemCondition** | **optional.String**| Filters the offer listings based on item condition. Possible values: New, Used, Collectible, Refurbished, Club. | 
 **offerType** | **optional.String**| Indicates whether to request pricing information for the seller&#x27;s B2C or B2B offers. Default is B2C. | 

### Return type

[**GetPricingResponse**](GetPricingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

