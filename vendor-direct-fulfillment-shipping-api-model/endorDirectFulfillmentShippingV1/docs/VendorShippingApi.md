# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPackingSlip**](VendorShippingApi.md#GetPackingSlip) | **Get** /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber} | 
[**GetPackingSlips**](VendorShippingApi.md#GetPackingSlips) | **Get** /vendor/directFulfillment/shipping/v1/packingSlips | 
[**SubmitShipmentConfirmations**](VendorShippingApi.md#SubmitShipmentConfirmations) | **Post** /vendor/directFulfillment/shipping/v1/shipmentConfirmations | 
[**SubmitShipmentStatusUpdates**](VendorShippingApi.md#SubmitShipmentStatusUpdates) | **Post** /vendor/directFulfillment/shipping/v1/shipmentStatusUpdates | 

# **GetPackingSlip**
> GetPackingSlipResponse GetPackingSlip(ctx, purchaseOrderNumber)


Returns a packing slip based on the purchaseOrderNumber that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purchaseOrderNumber** | **string**| The purchaseOrderNumber for the packing slip you want. | 

### Return type

[**GetPackingSlipResponse**](GetPackingSlipResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPackingSlips**
> GetPackingSlipListResponse GetPackingSlips(ctx, createdAfter, createdBefore, optional)


Returns a list of packing slips for the purchase orders that match the criteria specified. Date range to search must not be more than 7 days.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createdAfter** | **time.Time**| Packing slips that became available after this date and time will be included in the result. Must be in ISO-8601 date/time format. | 
  **createdBefore** | **time.Time**| Packing slips that became available before this date and time will be included in the result. Must be in ISO-8601 date/time format. | 
 **optional** | ***VendorShippingApiGetPackingSlipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorShippingApiGetPackingSlipsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipFromPartyId** | **optional.String**| The vendor warehouseId for order fulfillment. If not specified the result will contain orders for all warehouses. | 
 **limit** | **optional.Int32**| The limit to the number of records returned | 
 **sortOrder** | **optional.String**| Sort ASC or DESC by packing slip creation date. | [default to ASC]
 **nextToken** | **optional.String**| Used for pagination when there are more packing slips than the specified result size limit. The token value is returned in the previous API call. | 

### Return type

[**GetPackingSlipListResponse**](GetPackingSlipListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitShipmentConfirmations**
> SubmitShipmentConfirmationsResponse SubmitShipmentConfirmations(ctx, body)


Submits one or more shipment confirmations for vendor orders.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitShipmentConfirmationsRequest**](SubmitShipmentConfirmationsRequest.md)|  | 

### Return type

[**SubmitShipmentConfirmationsResponse**](SubmitShipmentConfirmationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitShipmentStatusUpdates**
> SubmitShipmentStatusUpdatesResponse SubmitShipmentStatusUpdates(ctx, body)


This API call is only to be used by Vendor-Own-Carrier (VOC) vendors. Calling this API will submit a shipment status update for the package that a vendor has shipped. It will provide the Amazon customer visibility on their order, when the package is outside of Amazon Network visibility.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitShipmentStatusUpdatesRequest**](SubmitShipmentStatusUpdatesRequest.md)|  | 

### Return type

[**SubmitShipmentStatusUpdatesResponse**](SubmitShipmentStatusUpdatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

