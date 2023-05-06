# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUploadDestinationForResource**](UploadsApi.md#CreateUploadDestinationForResource) | **Post** /uploads/2020-11-01/uploadDestinations/{resource} | 

# **CreateUploadDestinationForResource**
> CreateUploadDestinationResponse CreateUploadDestinationForResource(ctx, marketplaceIds, contentMD5, resource, optional)


Creates an upload destination, returning the information required to upload a file to the destination and to programmatically access the file.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceIds** | [**[]string**](string.md)| A list of marketplace identifiers. This specifies the marketplaces where the upload will be available. Only one marketplace can be specified. | 
  **contentMD5** | **string**| An MD5 hash of the content to be submitted to the upload destination. This value is used to determine if the data has been corrupted or tampered with during transit. | 
  **resource** | **string**| The resource for the upload destination that you are creating. For example, if you are creating an upload destination for the createLegalDisclosure operation of the Messaging API, the &#x60;{resource}&#x60; would be &#x60;/messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure&#x60;, and the entire path would be &#x60;/uploads/2020-11-01/uploadDestinations/messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure&#x60;. If you are creating an upload destination for an Aplus content document, the &#x60;{resource}&#x60; would be &#x60;aplus/2020-11-01/contentDocuments&#x60; and the path would be &#x60;/uploads/v1/uploadDestinations/aplus/2020-11-01/contentDocuments&#x60;. | 
 **optional** | ***UploadsApiCreateUploadDestinationForResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadsApiCreateUploadDestinationForResourceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **optional.String**| The content type of the file to be uploaded. | 

### Return type

[**CreateUploadDestinationResponse**](CreateUploadDestinationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

