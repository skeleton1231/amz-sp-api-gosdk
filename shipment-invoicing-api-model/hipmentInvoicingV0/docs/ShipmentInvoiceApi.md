# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoiceStatus**](ShipmentInvoiceApi.md#GetInvoiceStatus) | **Get** /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status | 
[**GetShipmentDetails**](ShipmentInvoiceApi.md#GetShipmentDetails) | **Get** /fba/outbound/brazil/v0/shipments/{shipmentId} | 
[**SubmitInvoice**](ShipmentInvoiceApi.md#SubmitInvoice) | **Post** /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice | 

# **GetInvoiceStatus**
> GetInvoiceStatusResponse GetInvoiceStatus(ctx, shipmentId)


Returns the invoice status for the shipment you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1.133 | 25 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| The shipment identifier for the shipment. | 

### Return type

[**GetInvoiceStatusResponse**](GetInvoiceStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShipmentDetails**
> GetShipmentDetailsResponse GetShipmentDetails(ctx, shipmentId)


Returns the shipment details required to issue an invoice for the specified shipment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1.133 | 25 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| The identifier for the shipment. Get this value from the FBAOutboundShipmentStatus notification. For information about subscribing to notifications, see the [Notifications API Use Case Guide](doc:notifications-api-v1-use-case-guide). | 

### Return type

[**GetShipmentDetailsResponse**](GetShipmentDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitInvoice**
> SubmitInvoiceResponse SubmitInvoice(ctx, body, shipmentId)


Submits a shipment invoice document for a given shipment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1.133 | 25 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitInvoiceRequest**](SubmitInvoiceRequest.md)|  | 
  **shipmentId** | **string**| The identifier for the shipment. | 

### Return type

[**SubmitInvoiceResponse**](SubmitInvoiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

