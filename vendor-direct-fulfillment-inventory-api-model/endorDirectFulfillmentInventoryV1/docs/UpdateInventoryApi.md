# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitInventoryUpdate**](UpdateInventoryApi.md#SubmitInventoryUpdate) | **Post** /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items | 

# **SubmitInventoryUpdate**
> SubmitInventoryUpdateResponse SubmitInventoryUpdate(ctx, body, warehouseId)


Submits inventory updates for the specified warehouse for either a partial or full feed of inventory items.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitInventoryUpdateRequest**](SubmitInventoryUpdateRequest.md)|  | 
  **warehouseId** | **string**| Identifier for the warehouse for which to update inventory. | 

### Return type

[**SubmitInventoryUpdateResponse**](SubmitInventoryUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

