# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContentDocument**](AplusContentApi.md#CreateContentDocument) | **Post** /aplus/2020-11-01/contentDocuments | 
[**GetContentDocument**](AplusContentApi.md#GetContentDocument) | **Get** /aplus/2020-11-01/contentDocuments/{contentReferenceKey} | 
[**ListContentDocumentAsinRelations**](AplusContentApi.md#ListContentDocumentAsinRelations) | **Get** /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins | 
[**PostContentDocumentApprovalSubmission**](AplusContentApi.md#PostContentDocumentApprovalSubmission) | **Post** /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/approvalSubmissions | 
[**PostContentDocumentAsinRelations**](AplusContentApi.md#PostContentDocumentAsinRelations) | **Post** /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins | 
[**PostContentDocumentSuspendSubmission**](AplusContentApi.md#PostContentDocumentSuspendSubmission) | **Post** /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/suspendSubmissions | 
[**SearchContentDocuments**](AplusContentApi.md#SearchContentDocuments) | **Get** /aplus/2020-11-01/contentDocuments | 
[**SearchContentPublishRecords**](AplusContentApi.md#SearchContentPublishRecords) | **Get** /aplus/2020-11-01/contentPublishRecords | 
[**UpdateContentDocument**](AplusContentApi.md#UpdateContentDocument) | **Post** /aplus/2020-11-01/contentDocuments/{contentReferenceKey} | 
[**ValidateContentDocumentAsinRelations**](AplusContentApi.md#ValidateContentDocumentAsinRelations) | **Post** /aplus/2020-11-01/contentAsinValidations | 

# **CreateContentDocument**
> PostContentDocumentResponse CreateContentDocument(ctx, body, marketplaceId)


Creates a new A+ Content document.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostContentDocumentRequest**](PostContentDocumentRequest.md)| The content document request details. | 
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 

### Return type

[**PostContentDocumentResponse**](PostContentDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentDocument**
> GetContentDocumentResponse GetContentDocument(ctx, contentReferenceKey, marketplaceId, includedDataSet)


Returns an A+ Content document, if available.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentReferenceKey** | **string**| The unique reference key for the A+ Content document. A content reference key cannot form a permalink and may change in the future. A content reference key is not guaranteed to match any A+ Content identifier. | 
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 
  **includedDataSet** | [**[]string**](string.md)| The set of A+ Content data types to include in the response. | 

### Return type

[**GetContentDocumentResponse**](GetContentDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContentDocumentAsinRelations**
> ListContentDocumentAsinRelationsResponse ListContentDocumentAsinRelations(ctx, contentReferenceKey, marketplaceId, optional)


Returns a list of ASINs related to the specified A+ Content document, if available. If you do not include the asinSet parameter, the operation returns all ASINs related to the content document.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentReferenceKey** | **string**| The unique reference key for the A+ Content document. A content reference key cannot form a permalink and may change in the future. A content reference key is not guaranteed to match any A+ Content identifier. | 
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 
 **optional** | ***AplusContentApiListContentDocumentAsinRelationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AplusContentApiListContentDocumentAsinRelationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includedDataSet** | [**optional.Interface of []string**](string.md)| The set of A+ Content data types to include in the response. If you do not include this parameter, the operation returns the related ASINs without metadata. | 
 **asinSet** | [**optional.Interface of []string**](string.md)| The set of ASINs. | 
 **pageToken** | **optional.String**| A page token from the nextPageToken response element returned by your previous call to this operation. nextPageToken is returned when the results of a call exceed the page size. To get the next page of results, call the operation and include pageToken as the only parameter. Specifying pageToken with any other parameter will cause the request to fail. When no nextPageToken value is returned there are no more pages to return. A pageToken value is not usable across different operations. | 

### Return type

[**ListContentDocumentAsinRelationsResponse**](ListContentDocumentAsinRelationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostContentDocumentApprovalSubmission**
> PostContentDocumentApprovalSubmissionResponse PostContentDocumentApprovalSubmission(ctx, contentReferenceKey, marketplaceId)


Submits an A+ Content document for review, approval, and publishing.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentReferenceKey** | **string**| The unique reference key for the A+ Content document. A content reference key cannot form a permalink and may change in the future. A content reference key is not guaranteed to match any A+ content identifier. | 
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 

### Return type

[**PostContentDocumentApprovalSubmissionResponse**](PostContentDocumentApprovalSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostContentDocumentAsinRelations**
> PostContentDocumentAsinRelationsResponse PostContentDocumentAsinRelations(ctx, body, contentReferenceKey, marketplaceId)


Replaces all ASINs related to the specified A+ Content document, if available. This may add or remove ASINs, depending on the current set of related ASINs. Removing an ASIN has the side effect of suspending the content document from that ASIN.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostContentDocumentAsinRelationsRequest**](PostContentDocumentAsinRelationsRequest.md)| The content document ASIN relations request details. | 
  **contentReferenceKey** | **string**| The unique reference key for the A+ Content document. A content reference key cannot form a permalink and may change in the future. A content reference key is not guaranteed to match any A+ content identifier. | 
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 

### Return type

[**PostContentDocumentAsinRelationsResponse**](PostContentDocumentAsinRelationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostContentDocumentSuspendSubmission**
> PostContentDocumentSuspendSubmissionResponse PostContentDocumentSuspendSubmission(ctx, contentReferenceKey, marketplaceId)


Submits a request to suspend visible A+ Content. This neither deletes the content document nor the ASIN relations.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentReferenceKey** | **string**| The unique reference key for the A+ Content document. A content reference key cannot form a permalink and may change in the future. A content reference key is not guaranteed to match any A+ content identifier. | 
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 

### Return type

[**PostContentDocumentSuspendSubmissionResponse**](PostContentDocumentSuspendSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchContentDocuments**
> SearchContentDocumentsResponse SearchContentDocuments(ctx, marketplaceId, optional)


Returns a list of all A+ Content documents assigned to a selling partner. This operation returns only the metadata of the A+ Content documents. Call the getContentDocument operation to get the actual contents of the A+ Content documents.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 
 **optional** | ***AplusContentApiSearchContentDocumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AplusContentApiSearchContentDocumentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A page token from the nextPageToken response element returned by your previous call to this operation. nextPageToken is returned when the results of a call exceed the page size. To get the next page of results, call the operation and include pageToken as the only parameter. Specifying pageToken with any other parameter will cause the request to fail. When no nextPageToken value is returned there are no more pages to return. A pageToken value is not usable across different operations. | 

### Return type

[**SearchContentDocumentsResponse**](SearchContentDocumentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchContentPublishRecords**
> SearchContentPublishRecordsResponse SearchContentPublishRecords(ctx, marketplaceId, asin, optional)


Searches for A+ Content publishing records, if available.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 
  **asin** | **string**| The Amazon Standard Identification Number (ASIN). | 
 **optional** | ***AplusContentApiSearchContentPublishRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AplusContentApiSearchContentPublishRecordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageToken** | **optional.String**| A page token from the nextPageToken response element returned by your previous call to this operation. nextPageToken is returned when the results of a call exceed the page size. To get the next page of results, call the operation and include pageToken as the only parameter. Specifying pageToken with any other parameter will cause the request to fail. When no nextPageToken value is returned there are no more pages to return. A pageToken value is not usable across different operations. | 

### Return type

[**SearchContentPublishRecordsResponse**](SearchContentPublishRecordsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContentDocument**
> PostContentDocumentResponse UpdateContentDocument(ctx, body, contentReferenceKey, marketplaceId)


Updates an existing A+ Content document.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostContentDocumentRequest**](PostContentDocumentRequest.md)| The content document request details. | 
  **contentReferenceKey** | **string**| The unique reference key for the A+ Content document. A content reference key cannot form a permalink and may change in the future. A content reference key is not guaranteed to match any A+ Content identifier. | 
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 

### Return type

[**PostContentDocumentResponse**](PostContentDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateContentDocumentAsinRelations**
> ValidateContentDocumentAsinRelationsResponse ValidateContentDocumentAsinRelations(ctx, body, marketplaceId, optional)


Checks if the A+ Content document is valid for use on a set of ASINs.  **Usage Plans:**  | Plan type | Rate (requests per second) | Burst | | ---- | ---- | ---- | |Default| 10 | 10 | |Selling partner specific| Variable | Variable |  The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see \"Usage Plans and Rate Limits\" in the Selling Partner API documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostContentDocumentRequest**](PostContentDocumentRequest.md)| The content document request details. | 
  **marketplaceId** | **string**| The identifier for the marketplace where the A+ Content is published. | 
 **optional** | ***AplusContentApiValidateContentDocumentAsinRelationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AplusContentApiValidateContentDocumentAsinRelationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asinSet** | [**optional.Interface of []string**](string.md)| The set of ASINs. | 

### Return type

[**ValidateContentDocumentAsinRelationsResponse**](ValidateContentDocumentAsinRelationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

