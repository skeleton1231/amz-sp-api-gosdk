# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefinitionsProductType**](DefinitionsApi.md#GetDefinitionsProductType) | **Get** /definitions/2020-09-01/productTypes/{productType} | 
[**SearchDefinitionsProductTypes**](DefinitionsApi.md#SearchDefinitionsProductTypes) | **Get** /definitions/2020-09-01/productTypes | 

# **GetDefinitionsProductType**
> ProductTypeDefinition GetDefinitionsProductType(ctx, productType, marketplaceIds, optional)


Retrieve an Amazon product type definition.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 5 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productType** | **string**| The Amazon product type name. | 
  **marketplaceIds** | [**[]string**](string.md)| A comma-delimited list of Amazon marketplace identifiers for the request. Note: This parameter is limited to one marketplaceId at this time. | 
 **optional** | ***DefinitionsApiGetDefinitionsProductTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinitionsApiGetDefinitionsProductTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sellerId** | **optional.String**| A selling partner identifier. When provided, seller-specific requirements and values are populated within the product type definition schema, such as brand names associated with the selling partner. | 
 **productTypeVersion** | **optional.String**| The version of the Amazon product type to retrieve. Defaults to \&quot;LATEST\&quot;,. Prerelease versions of product type definitions may be retrieved with \&quot;RELEASE_CANDIDATE\&quot;. If no prerelease version is currently available, the \&quot;LATEST\&quot; live version will be provided. | [default to LATEST]
 **requirements** | **optional.String**| The name of the requirements set to retrieve requirements for. | [default to LISTING]
 **requirementsEnforced** | **optional.String**| Identifies if the required attributes for a requirements set are enforced by the product type definition schema. Non-enforced requirements enable structural validation of individual attributes without all the required attributes being present (such as for partial updates). | [default to ENFORCED]
 **locale** | **optional.String**| Locale for retrieving display labels and other presentation details. Defaults to the default language of the first marketplace in the request. | [default to DEFAULT]

### Return type

[**ProductTypeDefinition**](ProductTypeDefinition.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchDefinitionsProductTypes**
> ProductTypeList SearchDefinitionsProductTypes(ctx, marketplaceIds, optional)


Search for and return a list of Amazon product types that have definitions available.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 5 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceIds** | [**[]string**](string.md)| A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **optional** | ***DefinitionsApiSearchDefinitionsProductTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinitionsApiSearchDefinitionsProductTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keywords** | [**optional.Interface of []string**](string.md)| A comma-delimited list of keywords to search product types by. | 

### Return type

[**ProductTypeList**](ProductTypeList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

