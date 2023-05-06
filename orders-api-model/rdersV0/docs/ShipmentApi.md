# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateShipmentStatus**](ShipmentApi.md#UpdateShipmentStatus) | **Post** /orders/v0/orders/{orderId}/shipment | 

# **UpdateShipmentStatus**
> UpdateShipmentStatus(ctx, body, orderId)


Update the shipment status for an order that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 15 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateShipmentStatusRequest**](UpdateShipmentStatusRequest.md)| The request body for the updateShipmentStatus operation. | 
  **orderId** | **string**| An Amazon-defined order identifier, in 3-7-7 format. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

