# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShipmentDetails**](VendorShippingApi.md#GetShipmentDetails) | **Get** /vendor/shipping/v1/shipments | 
[**GetShipmentLabels**](VendorShippingApi.md#GetShipmentLabels) | **Get** /vendor/shipping/v1/transportLabels | 
[**SubmitShipmentConfirmations**](VendorShippingApi.md#SubmitShipmentConfirmations) | **Post** /vendor/shipping/v1/shipmentConfirmations | 
[**SubmitShipments**](VendorShippingApi.md#SubmitShipments) | **Post** /vendor/shipping/v1/shipments | 

# **GetShipmentDetails**
> GetShipmentDetailsResponse GetShipmentDetails(ctx, optional)


Returns the Details about Shipment, Carrier Details,  status of the shipment, container details and other details related to shipment based on the filter parameters value that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VendorShippingApiGetShipmentDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorShippingApiGetShipmentDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| The limit to the number of records returned. Default value is 50 records. | 
 **sortOrder** | **optional.String**| Sort in ascending or descending order by purchase order creation date. | 
 **nextToken** | **optional.String**| Used for pagination when there are more shipments than the specified result size limit. | 
 **createdAfter** | **optional.Time**| Get Shipment Details that became available after this timestamp will be included in the result. Must be in ISO-8601 date/time format. | 
 **createdBefore** | **optional.Time**| Get Shipment Details that became available before this timestamp will be included in the result. Must be in ISO-8601 date/time format. | 
 **shipmentConfirmedBefore** | **optional.Time**| Get Shipment Details by passing Shipment confirmed create Date Before. Must be in ISO-8601 date/time format. | 
 **shipmentConfirmedAfter** | **optional.Time**| Get Shipment Details by passing Shipment confirmed create Date After. Must be in ISO-8601 date/time format. | 
 **packageLabelCreatedBefore** | **optional.Time**| Get Shipment Details by passing Package label create Date by buyer. Must be in ISO-8601 date/time format. | 
 **packageLabelCreatedAfter** | **optional.Time**| Get Shipment Details by passing Package label create Date After by buyer. Must be in ISO-8601 date/time format. | 
 **shippedBefore** | **optional.Time**| Get Shipment Details by passing Shipped Date Before. Must be in ISO-8601 date/time format. | 
 **shippedAfter** | **optional.Time**| Get Shipment Details by passing Shipped Date After. Must be in ISO-8601 date/time format. | 
 **estimatedDeliveryBefore** | **optional.Time**| Get Shipment Details by passing Estimated Delivery Date Before. Must be in ISO-8601 date/time format. | 
 **estimatedDeliveryAfter** | **optional.Time**| Get Shipment Details by passing Estimated Delivery Date Before. Must be in ISO-8601 date/time format. | 
 **shipmentDeliveryBefore** | **optional.Time**| Get Shipment Details by passing Shipment Delivery Date Before. Must be in ISO-8601 date/time format. | 
 **shipmentDeliveryAfter** | **optional.Time**| Get Shipment Details by passing Shipment Delivery Date After. Must be in ISO-8601 date/time format. | 
 **requestedPickUpBefore** | **optional.Time**| Get Shipment Details by passing Before Requested pickup date. Must be in ISO-8601 date/time format. | 
 **requestedPickUpAfter** | **optional.Time**| Get Shipment Details by passing After Requested pickup date. Must be in ISO-8601 date/time format. | 
 **scheduledPickUpBefore** | **optional.Time**| Get Shipment Details by passing Before scheduled pickup date. Must be in ISO-8601 date/time format. | 
 **scheduledPickUpAfter** | **optional.Time**| Get Shipment Details by passing After Scheduled pickup date. Must be in ISO-8601 date/time format. | 
 **currentShipmentStatus** | **optional.String**| Get Shipment Details by passing Current shipment status. | 
 **vendorShipmentIdentifier** | **optional.String**| Get Shipment Details by passing Vendor Shipment ID | 
 **buyerReferenceNumber** | **optional.String**| Get Shipment Details by passing buyer Reference ID | 
 **buyerWarehouseCode** | **optional.String**| Get Shipping Details based on buyer warehouse code. This value should be same as &#x27;shipToParty.partyId&#x27; in the Shipment. | 
 **sellerWarehouseCode** | **optional.String**| Get Shipping Details based on vendor warehouse code. This value should be same as &#x27;sellingParty.partyId&#x27; in the Shipment. | 

### Return type

[**GetShipmentDetailsResponse**](GetShipmentDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShipmentLabels**
> GetShipmentLabels GetShipmentLabels(ctx, optional)


Returns transport Labels based on the filters that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VendorShippingApiGetShipmentLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorShippingApiGetShipmentLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| The limit to the number of records returned. Default value is 50 records. | 
 **sortOrder** | **optional.String**| Sort in ascending or descending order by transport label creation date. | 
 **nextToken** | **optional.String**| Used for pagination when there are more transport label than the specified result size limit. | 
 **labelCreatedAfter** | **optional.Time**| transport Labels that became available after this timestamp will be included in the result. Must be in ISO-8601 date/time format. | 
 **labelcreatedBefore** | **optional.Time**| transport Labels that became available before this timestamp will be included in the result. Must be in ISO-8601 date/time format. | 
 **buyerReferenceNumber** | **optional.String**| Get transport labels by passing Buyer Reference Number to retreive the corresponding transport label. | 
 **vendorShipmentIdentifier** | **optional.String**| Get transport labels by passing Vendor Shipment ID to retreive the corresponding transport label. | 
 **sellerWarehouseCode** | **optional.String**| Get Shipping labels based Vendor Warehouse code. This value should be same as &#x27;shipFromParty.partyId&#x27; in the Shipment. | 

### Return type

[**GetShipmentLabels**](GetShipmentLabels.md)

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

# **SubmitShipments**
> SubmitShipmentConfirmationsResponse SubmitShipments(ctx, body)


Submits one or more shipment request for vendor Orders.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitShipments**](SubmitShipments.md)|  | 

### Return type

[**SubmitShipmentConfirmationsResponse**](SubmitShipmentConfirmationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

