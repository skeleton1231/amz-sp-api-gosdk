# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomerInvoice**](CustomerInvoicesApi.md#GetCustomerInvoice) | **Get** /vendor/directFulfillment/shipping/2021-12-28/customerInvoices/{purchaseOrderNumber} | 
[**GetCustomerInvoices**](CustomerInvoicesApi.md#GetCustomerInvoices) | **Get** /vendor/directFulfillment/shipping/2021-12-28/customerInvoices | 

# **GetCustomerInvoice**
> CustomerInvoice GetCustomerInvoice(ctx, purchaseOrderNumber)


Returns a customer invoice based on the purchaseOrderNumber that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purchaseOrderNumber** | **string**| Purchase order number of the shipment for which to return the invoice. | 

### Return type

[**CustomerInvoice**](CustomerInvoice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerInvoices**
> CustomerInvoiceList GetCustomerInvoices(ctx, createdAfter, createdBefore, optional)


Returns a list of customer invoices created during a time frame that you specify. You define the time frame using the createdAfter and createdBefore parameters. You must use both of these parameters. The date range to search must be no more than 7 days.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createdAfter** | **time.Time**| Orders that became available after this date and time will be included in the result. Must be in ISO-8601 date/time format. | 
  **createdBefore** | **time.Time**| Orders that became available before this date and time will be included in the result. Must be in ISO-8601 date/time format. | 
 **optional** | ***CustomerInvoicesApiGetCustomerInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerInvoicesApiGetCustomerInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipFromPartyId** | **optional.String**| The vendor warehouseId for order fulfillment. If not specified, the result will contain orders for all warehouses. | 
 **limit** | **optional.Int32**| The limit to the number of records returned | 
 **sortOrder** | **optional.String**| Sort ASC or DESC by order creation date. | 
 **nextToken** | **optional.String**| Used for pagination when there are more orders than the specified result size limit. The token value is returned in the previous API call. | 

### Return type

[**CustomerInvoiceList**](CustomerInvoiceList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, payload

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

