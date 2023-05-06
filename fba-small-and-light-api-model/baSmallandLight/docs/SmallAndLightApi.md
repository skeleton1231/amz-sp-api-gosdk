# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSmallAndLightEnrollmentBySellerSKU**](SmallAndLightApi.md#DeleteSmallAndLightEnrollmentBySellerSKU) | **Delete** /fba/smallAndLight/v1/enrollments/{sellerSKU} | 
[**GetSmallAndLightEligibilityBySellerSKU**](SmallAndLightApi.md#GetSmallAndLightEligibilityBySellerSKU) | **Get** /fba/smallAndLight/v1/eligibilities/{sellerSKU} | 
[**GetSmallAndLightEnrollmentBySellerSKU**](SmallAndLightApi.md#GetSmallAndLightEnrollmentBySellerSKU) | **Get** /fba/smallAndLight/v1/enrollments/{sellerSKU} | 
[**GetSmallAndLightFeePreview**](SmallAndLightApi.md#GetSmallAndLightFeePreview) | **Post** /fba/smallAndLight/v1/feePreviews | 
[**PutSmallAndLightEnrollmentBySellerSKU**](SmallAndLightApi.md#PutSmallAndLightEnrollmentBySellerSKU) | **Put** /fba/smallAndLight/v1/enrollments/{sellerSKU} | 

# **DeleteSmallAndLightEnrollmentBySellerSKU**
> DeleteSmallAndLightEnrollmentBySellerSKU(ctx, sellerSKU, marketplaceIds)


Removes the item indicated by the specified seller SKU from the Small and Light program in the specified marketplace. If the item is not eligible for disenrollment, the ineligibility reasons are returned.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sellerSKU** | **string**| The seller SKU that identifies the item. | 
  **marketplaceIds** | [**[]string**](string.md)| The marketplace in which to remove the item from the Small and Light program. Note: Accepts a single marketplace only. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmallAndLightEligibilityBySellerSKU**
> SmallAndLightEligibility GetSmallAndLightEligibilityBySellerSKU(ctx, sellerSKU, marketplaceIds)


Returns the Small and Light program eligibility status of the item indicated by the specified seller SKU in the specified marketplace. If the item is not eligible, the ineligibility reasons are returned. **Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sellerSKU** | **string**| The seller SKU that identifies the item. | 
  **marketplaceIds** | [**[]string**](string.md)| The marketplace for which the eligibility status is retrieved. NOTE: Accepts a single marketplace only. | 

### Return type

[**SmallAndLightEligibility**](SmallAndLightEligibility.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmallAndLightEnrollmentBySellerSKU**
> SmallAndLightEnrollment GetSmallAndLightEnrollmentBySellerSKU(ctx, sellerSKU, marketplaceIds)


Returns the Small and Light enrollment status for the item indicated by the specified seller SKU in the specified marketplace.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sellerSKU** | **string**| The seller SKU that identifies the item. | 
  **marketplaceIds** | [**[]string**](string.md)| The marketplace for which the enrollment status is retrieved. Note: Accepts a single marketplace only. | 

### Return type

[**SmallAndLightEnrollment**](SmallAndLightEnrollment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmallAndLightFeePreview**
> SmallAndLightFeePreviews GetSmallAndLightFeePreview(ctx, body)


Returns the Small and Light fee estimates for the specified items. You must include a marketplaceId parameter to retrieve the proper fee estimates for items to be sold in that marketplace. The ordering of items in the response will mirror the order of the items in the request. Duplicate ASIN/price combinations are removed.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 3 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmallAndLightFeePreviewRequest**](SmallAndLightFeePreviewRequest.md)|  | 

### Return type

[**SmallAndLightFeePreviews**](SmallAndLightFeePreviews.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSmallAndLightEnrollmentBySellerSKU**
> SmallAndLightEnrollment PutSmallAndLightEnrollmentBySellerSKU(ctx, sellerSKU, marketplaceIds)


Enrolls the item indicated by the specified seller SKU in the Small and Light program in the specified marketplace. If the item is not eligible, the ineligibility reasons are returned.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sellerSKU** | **string**| The seller SKU that identifies the item. | 
  **marketplaceIds** | [**[]string**](string.md)| The marketplace in which to enroll the item. Note: Accepts a single marketplace only. | 

### Return type

[**SmallAndLightEnrollment**](SmallAndLightEnrollment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

