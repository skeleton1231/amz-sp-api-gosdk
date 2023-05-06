# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmPreorder**](FbaInboundApi.md#ConfirmPreorder) | **Put** /fba/inbound/v0/shipments/{shipmentId}/preorder/confirm | 
[**ConfirmTransport**](FbaInboundApi.md#ConfirmTransport) | **Post** /fba/inbound/v0/shipments/{shipmentId}/transport/confirm | 
[**CreateInboundShipment**](FbaInboundApi.md#CreateInboundShipment) | **Post** /fba/inbound/v0/shipments/{shipmentId} | 
[**CreateInboundShipmentPlan**](FbaInboundApi.md#CreateInboundShipmentPlan) | **Post** /fba/inbound/v0/plans | 
[**EstimateTransport**](FbaInboundApi.md#EstimateTransport) | **Post** /fba/inbound/v0/shipments/{shipmentId}/transport/estimate | 
[**GetBillOfLading**](FbaInboundApi.md#GetBillOfLading) | **Get** /fba/inbound/v0/shipments/{shipmentId}/billOfLading | 
[**GetInboundGuidance**](FbaInboundApi.md#GetInboundGuidance) | **Get** /fba/inbound/v0/itemsGuidance | 
[**GetLabels**](FbaInboundApi.md#GetLabels) | **Get** /fba/inbound/v0/shipments/{shipmentId}/labels | 
[**GetPreorderInfo**](FbaInboundApi.md#GetPreorderInfo) | **Get** /fba/inbound/v0/shipments/{shipmentId}/preorder | 
[**GetPrepInstructions**](FbaInboundApi.md#GetPrepInstructions) | **Get** /fba/inbound/v0/prepInstructions | 
[**GetShipmentItems**](FbaInboundApi.md#GetShipmentItems) | **Get** /fba/inbound/v0/shipmentItems | 
[**GetShipmentItemsByShipmentId**](FbaInboundApi.md#GetShipmentItemsByShipmentId) | **Get** /fba/inbound/v0/shipments/{shipmentId}/items | 
[**GetShipments**](FbaInboundApi.md#GetShipments) | **Get** /fba/inbound/v0/shipments | 
[**GetTransportDetails**](FbaInboundApi.md#GetTransportDetails) | **Get** /fba/inbound/v0/shipments/{shipmentId}/transport | 
[**PutTransportDetails**](FbaInboundApi.md#PutTransportDetails) | **Put** /fba/inbound/v0/shipments/{shipmentId}/transport | 
[**UpdateInboundShipment**](FbaInboundApi.md#UpdateInboundShipment) | **Put** /fba/inbound/v0/shipments/{shipmentId} | 
[**VoidTransport**](FbaInboundApi.md#VoidTransport) | **Post** /fba/inbound/v0/shipments/{shipmentId}/transport/void | 

# **ConfirmPreorder**
> ConfirmPreorderResponse ConfirmPreorder(ctx, shipmentId, needByDate, marketplaceId)


Returns information needed to confirm a shipment for pre-order. Call this operation after calling the getPreorderInfo operation to get the NeedByDate value and other pre-order information about the shipment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 
  **needByDate** | **string**| Date that the shipment must arrive at the Amazon fulfillment center to avoid delivery promise breaks for pre-ordered items. Must be in YYYY-MM-DD format. The response to the getPreorderInfo operation returns this value. | 
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace the shipment is tied to. | 

### Return type

[**ConfirmPreorderResponse**](ConfirmPreorderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfirmTransport**
> ConfirmTransportResponse ConfirmTransport(ctx, shipmentId)


Confirms that the seller accepts the Amazon-partnered shipping estimate, agrees to allow Amazon to charge their account for the shipping cost, and requests that the Amazon-partnered carrier ship the inbound shipment.  Prior to calling the confirmTransport operation, you should call the getTransportDetails operation to get the Amazon-partnered shipping estimate.  Important: After confirming the transportation request, if the seller decides that they do not want the Amazon-partnered carrier to ship the inbound shipment, you can call the voidTransport operation to cancel the transportation request. Note that for a Small Parcel shipment, the seller has 24 hours after confirming a transportation request to void the transportation request. For a Less Than Truckload/Full Truckload (LTL/FTL) shipment, the seller has one hour after confirming a transportation request to void it. After the grace period has expired the seller's account will be charged for the shipping cost.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 

### Return type

[**ConfirmTransportResponse**](ConfirmTransportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInboundShipment**
> InboundShipmentResponse CreateInboundShipment(ctx, body, shipmentId)


Returns a new inbound shipment based on the specified shipmentId that was returned by the createInboundShipmentPlan operation.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundShipmentRequest**](InboundShipmentRequest.md)|  | 
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 

### Return type

[**InboundShipmentResponse**](InboundShipmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInboundShipmentPlan**
> CreateInboundShipmentPlanResponse CreateInboundShipmentPlan(ctx, body)


Returns one or more inbound shipment plans, which provide the information you need to create one or more inbound shipments for a set of items that you specify. Multiple inbound shipment plans might be required so that items can be optimally placed in Amazon's fulfillment network—for example, positioning inventory closer to the customer. Alternatively, two inbound shipment plans might be created with the same Amazon fulfillment center destination if the two shipment plans require different processing—for example, items that require labels must be shipped separately from stickerless, commingled inventory.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateInboundShipmentPlanRequest**](CreateInboundShipmentPlanRequest.md)|  | 

### Return type

[**CreateInboundShipmentPlanResponse**](CreateInboundShipmentPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstimateTransport**
> EstimateTransportResponse EstimateTransport(ctx, shipmentId)


Initiates the process of estimating the shipping cost for an inbound shipment by an Amazon-partnered carrier.  Prior to calling the estimateTransport operation, you must call the putTransportDetails operation to provide Amazon with the transportation information for the inbound shipment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 

### Return type

[**EstimateTransportResponse**](EstimateTransportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBillOfLading**
> GetBillOfLadingResponse GetBillOfLading(ctx, shipmentId)


Returns a bill of lading for a Less Than Truckload/Full Truckload (LTL/FTL) shipment. The getBillOfLading operation returns PDF document data for printing a bill of lading for an Amazon-partnered Less Than Truckload/Full Truckload (LTL/FTL) inbound shipment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 

### Return type

[**GetBillOfLadingResponse**](GetBillOfLadingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInboundGuidance**
> GetInboundGuidanceResponse GetInboundGuidance(ctx, marketplaceId, optional)


Returns information that lets a seller know if Amazon recommends sending an item to a given marketplace. In some cases, Amazon provides guidance for why a given SellerSKU or ASIN is not recommended for shipment to Amazon's fulfillment network. Sellers may still ship items that are not recommended, at their discretion.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace where the product would be stored. | 
 **optional** | ***FbaInboundApiGetInboundGuidanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FbaInboundApiGetInboundGuidanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sellerSKUList** | [**optional.Interface of []string**](string.md)| A list of SellerSKU values. Used to identify items for which you want inbound guidance for shipment to Amazon&#x27;s fulfillment network. Note: SellerSKU is qualified by the SellerId, which is included with every Selling Partner API operation that you submit. If you specify a SellerSKU that identifies a variation parent ASIN, this operation returns an error. A variation parent ASIN represents a generic product that cannot be sold. Variation child ASINs represent products that have specific characteristics (such as size and color) and can be sold.  | 
 **aSINList** | [**optional.Interface of []string**](string.md)| A list of ASIN values. Used to identify items for which you want inbound guidance for shipment to Amazon&#x27;s fulfillment network. Note: If you specify a ASIN that identifies a variation parent ASIN, this operation returns an error. A variation parent ASIN represents a generic product that cannot be sold. Variation child ASINs represent products that have specific characteristics (such as size and color) and can be sold. | 

### Return type

[**GetInboundGuidanceResponse**](GetInboundGuidanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLabels**
> GetLabelsResponse GetLabels(ctx, shipmentId, pageType, labelType, optional)


Returns package/pallet labels for faster and more accurate shipment processing at the Amazon fulfillment center.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 
  **pageType** | **string**| The page type to use to print the labels. Submitting a PageType value that is not supported in your marketplace returns an error. | 
  **labelType** | **string**| The type of labels requested.  | 
 **optional** | ***FbaInboundApiGetLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FbaInboundApiGetLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **numberOfPackages** | **optional.Int32**| The number of packages in the shipment. | 
 **packageLabelsToPrint** | [**optional.Interface of []string**](string.md)| A list of identifiers that specify packages for which you want package labels printed.  Must match CartonId values previously passed using the FBA Inbound Shipment Carton Information Feed. If not, the operation returns the IncorrectPackageIdentifier error code. | 
 **numberOfPallets** | **optional.Int32**| The number of pallets in the shipment. This returns four identical labels for each pallet. | 
 **pageSize** | **optional.Int32**| The page size for paginating through the total packages&#x27; labels. This is a required parameter for Non-Partnered LTL Shipments. Max value:1000. | 
 **pageStartIndex** | **optional.Int32**| The page start index for paginating through the total packages&#x27; labels. This is a required parameter for Non-Partnered LTL Shipments. | 

### Return type

[**GetLabelsResponse**](GetLabelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreorderInfo**
> GetPreorderInfoResponse GetPreorderInfo(ctx, shipmentId, marketplaceId)


Returns pre-order information, including dates, that a seller needs before confirming a shipment for pre-order.   **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace the shipment is tied to. | 

### Return type

[**GetPreorderInfoResponse**](GetPreorderInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrepInstructions**
> GetPrepInstructionsResponse GetPrepInstructions(ctx, shipToCountryCode, optional)


Returns labeling requirements and item preparation instructions to help prepare items for shipment to Amazon's fulfillment network.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipToCountryCode** | **string**| The country code of the country to which the items will be shipped. Note that labeling requirements and item preparation instructions can vary by country. | 
 **optional** | ***FbaInboundApiGetPrepInstructionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FbaInboundApiGetPrepInstructionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sellerSKUList** | [**optional.Interface of []string**](string.md)| A list of SellerSKU values. Used to identify items for which you want labeling requirements and item preparation instructions for shipment to Amazon&#x27;s fulfillment network. The SellerSKU is qualified by the Seller ID, which is included with every call to the Seller Partner API.  Note: Include seller SKUs that you have used to list items on Amazon&#x27;s retail website. If you include a seller SKU that you have never used to list an item on Amazon&#x27;s retail website, the seller SKU is returned in the InvalidSKUList property in the response. | 
 **aSINList** | [**optional.Interface of []string**](string.md)| A list of ASIN values. Used to identify items for which you want item preparation instructions to help with item sourcing decisions.  Note: ASINs must be included in the product catalog for at least one of the marketplaces that the seller  participates in. Any ASIN that is not included in the product catalog for at least one of the marketplaces that the seller participates in is returned in the InvalidASINList property in the response. You can find out which marketplaces a seller participates in by calling the getMarketplaceParticipations operation in the Selling Partner API for Sellers. | 

### Return type

[**GetPrepInstructionsResponse**](GetPrepInstructionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShipmentItems**
> GetShipmentItemsResponse GetShipmentItems(ctx, queryType, marketplaceId, optional)


Returns a list of items in a specified inbound shipment, or a list of items that were updated within a specified time frame.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryType** | **string**| Indicates whether items are returned using a date range (by providing the LastUpdatedAfter and LastUpdatedBefore parameters), or using NextToken, which continues returning items specified in a previous request. | 
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace where the product would be stored. | 
 **optional** | ***FbaInboundApiGetShipmentItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FbaInboundApiGetShipmentItemsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lastUpdatedAfter** | **optional.Time**| A date used for selecting inbound shipment items that were last updated after (or at) a specified time. The selection includes updates made by Amazon and by the seller. | 
 **lastUpdatedBefore** | **optional.Time**| A date used for selecting inbound shipment items that were last updated before (or at) a specified time. The selection includes updates made by Amazon and by the seller. | 
 **nextToken** | **optional.String**| A string token returned in the response to your previous request. | 

### Return type

[**GetShipmentItemsResponse**](GetShipmentItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShipmentItemsByShipmentId**
> GetShipmentItemsResponse GetShipmentItemsByShipmentId(ctx, shipmentId, marketplaceId)


Returns a list of items in a specified inbound shipment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| A shipment identifier used for selecting items in a specific inbound shipment. | 
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace where the product would be stored. | 

### Return type

[**GetShipmentItemsResponse**](GetShipmentItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShipments**
> GetShipmentsResponse GetShipments(ctx, queryType, marketplaceId, optional)


Returns a list of inbound shipments based on criteria that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryType** | **string**| Indicates whether shipments are returned using shipment information (by providing the ShipmentStatusList or ShipmentIdList parameters), using a date range (by providing the LastUpdatedAfter and LastUpdatedBefore parameters), or by using NextToken to continue returning items specified in a previous request. | 
  **marketplaceId** | **string**| A marketplace identifier. Specifies the marketplace where the product would be stored. | 
 **optional** | ***FbaInboundApiGetShipmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FbaInboundApiGetShipmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipmentStatusList** | [**optional.Interface of []string**](string.md)| A list of ShipmentStatus values. Used to select shipments with a current status that matches the status values that you specify. | 
 **shipmentIdList** | [**optional.Interface of []string**](string.md)| A list of shipment IDs used to select the shipments that you want. If both ShipmentStatusList and ShipmentIdList are specified, only shipments that match both parameters are returned. | 
 **lastUpdatedAfter** | **optional.Time**| A date used for selecting inbound shipments that were last updated after (or at) a specified time. The selection includes updates made by Amazon and by the seller. | 
 **lastUpdatedBefore** | **optional.Time**| A date used for selecting inbound shipments that were last updated before (or at) a specified time. The selection includes updates made by Amazon and by the seller. | 
 **nextToken** | **optional.String**| A string token returned in the response to your previous request. | 

### Return type

[**GetShipmentsResponse**](GetShipmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportDetails**
> GetTransportDetailsResponse GetTransportDetails(ctx, shipmentId)


Returns current transportation information about an inbound shipment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 

### Return type

[**GetTransportDetailsResponse**](GetTransportDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTransportDetails**
> PutTransportDetailsResponse PutTransportDetails(ctx, body, shipmentId)


Sends transportation information to Amazon about an inbound shipment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutTransportDetailsRequest**](PutTransportDetailsRequest.md)|  | 
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 

### Return type

[**PutTransportDetailsResponse**](PutTransportDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInboundShipment**
> InboundShipmentResponse UpdateInboundShipment(ctx, body, shipmentId)


Updates or removes items from the inbound shipment identified by the specified shipment identifier. Adding new items is not supported.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundShipmentRequest**](InboundShipmentRequest.md)|  | 
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 

### Return type

[**InboundShipmentResponse**](InboundShipmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VoidTransport**
> VoidTransportResponse VoidTransport(ctx, shipmentId)


Cancels a previously-confirmed request to ship an inbound shipment using an Amazon-partnered carrier.  To be successful, you must call this operation before the VoidDeadline date that is returned by the getTransportDetails operation.  Important: The VoidDeadline date is 24 hours after you confirm a Small Parcel shipment transportation request or one hour after you confirm a Less Than Truckload/Full Truckload (LTL/FTL) shipment transportation request. After the void deadline passes, your account will be charged for the shipping cost.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 2 | 30 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentId** | **string**| A shipment identifier originally returned by the createInboundShipmentPlan operation. | 

### Return type

[**VoidTransportResponse**](VoidTransportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

