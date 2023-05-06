# Go API client for swagger

The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2020-07-01
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://sellercentral.amazon.com/gp/mws/contactus.html](https://sellercentral.amazon.com/gp/mws/contactus.html)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FbaOutboundApi* | [**CancelFulfillmentOrder**](docs/FbaOutboundApi.md#cancelfulfillmentorder) | **Put** /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel | 
*FbaOutboundApi* | [**CreateFulfillmentOrder**](docs/FbaOutboundApi.md#createfulfillmentorder) | **Post** /fba/outbound/2020-07-01/fulfillmentOrders | 
*FbaOutboundApi* | [**CreateFulfillmentReturn**](docs/FbaOutboundApi.md#createfulfillmentreturn) | **Put** /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return | 
*FbaOutboundApi* | [**GetFeatureInventory**](docs/FbaOutboundApi.md#getfeatureinventory) | **Get** /fba/outbound/2020-07-01/features/inventory/{featureName} | 
*FbaOutboundApi* | [**GetFeatureSKU**](docs/FbaOutboundApi.md#getfeaturesku) | **Get** /fba/outbound/2020-07-01/features/inventory/{featureName}/{sellerSku} | 
*FbaOutboundApi* | [**GetFeatures**](docs/FbaOutboundApi.md#getfeatures) | **Get** /fba/outbound/2020-07-01/features | 
*FbaOutboundApi* | [**GetFulfillmentOrder**](docs/FbaOutboundApi.md#getfulfillmentorder) | **Get** /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId} | 
*FbaOutboundApi* | [**GetFulfillmentPreview**](docs/FbaOutboundApi.md#getfulfillmentpreview) | **Post** /fba/outbound/2020-07-01/fulfillmentOrders/preview | 
*FbaOutboundApi* | [**GetPackageTrackingDetails**](docs/FbaOutboundApi.md#getpackagetrackingdetails) | **Get** /fba/outbound/2020-07-01/tracking | 
*FbaOutboundApi* | [**ListAllFulfillmentOrders**](docs/FbaOutboundApi.md#listallfulfillmentorders) | **Get** /fba/outbound/2020-07-01/fulfillmentOrders | 
*FbaOutboundApi* | [**ListReturnReasonCodes**](docs/FbaOutboundApi.md#listreturnreasoncodes) | **Get** /fba/outbound/2020-07-01/returnReasonCodes | 
*FbaOutboundApi* | [**SubmitFulfillmentOrderStatusUpdate**](docs/FbaOutboundApi.md#submitfulfillmentorderstatusupdate) | **Put** /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/status | 
*FbaOutboundApi* | [**UpdateFulfillmentOrder**](docs/FbaOutboundApi.md#updatefulfillmentorder) | **Put** /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId} | 

## Documentation For Models

 - [AdditionalLocationInfo](docs/AdditionalLocationInfo.md)
 - [Address](docs/Address.md)
 - [CancelFulfillmentOrderResponse](docs/CancelFulfillmentOrderResponse.md)
 - [CodSettings](docs/CodSettings.md)
 - [CreateFulfillmentOrderItem](docs/CreateFulfillmentOrderItem.md)
 - [CreateFulfillmentOrderRequest](docs/CreateFulfillmentOrderRequest.md)
 - [CreateFulfillmentOrderResponse](docs/CreateFulfillmentOrderResponse.md)
 - [CreateFulfillmentReturnRequest](docs/CreateFulfillmentReturnRequest.md)
 - [CreateFulfillmentReturnResponse](docs/CreateFulfillmentReturnResponse.md)
 - [CreateFulfillmentReturnResult](docs/CreateFulfillmentReturnResult.md)
 - [CreateReturnItem](docs/CreateReturnItem.md)
 - [CurrentStatus](docs/CurrentStatus.md)
 - [DeliveryWindow](docs/DeliveryWindow.md)
 - [EventCode](docs/EventCode.md)
 - [Feature](docs/Feature.md)
 - [FeatureSettings](docs/FeatureSettings.md)
 - [FeatureSku](docs/FeatureSku.md)
 - [Fee](docs/Fee.md)
 - [FulfillmentAction](docs/FulfillmentAction.md)
 - [FulfillmentOrder](docs/FulfillmentOrder.md)
 - [FulfillmentOrderItem](docs/FulfillmentOrderItem.md)
 - [FulfillmentOrderStatus](docs/FulfillmentOrderStatus.md)
 - [FulfillmentPolicy](docs/FulfillmentPolicy.md)
 - [FulfillmentPreview](docs/FulfillmentPreview.md)
 - [FulfillmentPreviewItem](docs/FulfillmentPreviewItem.md)
 - [FulfillmentPreviewShipment](docs/FulfillmentPreviewShipment.md)
 - [FulfillmentReturnItemStatus](docs/FulfillmentReturnItemStatus.md)
 - [FulfillmentShipment](docs/FulfillmentShipment.md)
 - [FulfillmentShipmentItem](docs/FulfillmentShipmentItem.md)
 - [FulfillmentShipmentPackage](docs/FulfillmentShipmentPackage.md)
 - [GetFeatureInventoryResponse](docs/GetFeatureInventoryResponse.md)
 - [GetFeatureInventoryResult](docs/GetFeatureInventoryResult.md)
 - [GetFeatureSkuResponse](docs/GetFeatureSkuResponse.md)
 - [GetFeatureSkuResult](docs/GetFeatureSkuResult.md)
 - [GetFeaturesResponse](docs/GetFeaturesResponse.md)
 - [GetFeaturesResult](docs/GetFeaturesResult.md)
 - [GetFulfillmentOrderResponse](docs/GetFulfillmentOrderResponse.md)
 - [GetFulfillmentOrderResult](docs/GetFulfillmentOrderResult.md)
 - [GetFulfillmentPreviewItem](docs/GetFulfillmentPreviewItem.md)
 - [GetFulfillmentPreviewRequest](docs/GetFulfillmentPreviewRequest.md)
 - [GetFulfillmentPreviewResponse](docs/GetFulfillmentPreviewResponse.md)
 - [GetFulfillmentPreviewResult](docs/GetFulfillmentPreviewResult.md)
 - [GetPackageTrackingDetailsResponse](docs/GetPackageTrackingDetailsResponse.md)
 - [InvalidItemReason](docs/InvalidItemReason.md)
 - [InvalidItemReasonCode](docs/InvalidItemReasonCode.md)
 - [InvalidReturnItem](docs/InvalidReturnItem.md)
 - [ListAllFulfillmentOrdersResponse](docs/ListAllFulfillmentOrdersResponse.md)
 - [ListAllFulfillmentOrdersResult](docs/ListAllFulfillmentOrdersResult.md)
 - [ListReturnReasonCodesResponse](docs/ListReturnReasonCodesResponse.md)
 - [ListReturnReasonCodesResult](docs/ListReturnReasonCodesResult.md)
 - [ModelError](docs/ModelError.md)
 - [Money](docs/Money.md)
 - [PackageTrackingDetails](docs/PackageTrackingDetails.md)
 - [ReasonCodeDetails](docs/ReasonCodeDetails.md)
 - [ReturnAuthorization](docs/ReturnAuthorization.md)
 - [ReturnItem](docs/ReturnItem.md)
 - [ReturnItemDisposition](docs/ReturnItemDisposition.md)
 - [ScheduledDeliveryInfo](docs/ScheduledDeliveryInfo.md)
 - [ShippingSpeedCategory](docs/ShippingSpeedCategory.md)
 - [SubmitFulfillmentOrderStatusUpdateRequest](docs/SubmitFulfillmentOrderStatusUpdateRequest.md)
 - [SubmitFulfillmentOrderStatusUpdateResponse](docs/SubmitFulfillmentOrderStatusUpdateResponse.md)
 - [TrackingAddress](docs/TrackingAddress.md)
 - [TrackingEvent](docs/TrackingEvent.md)
 - [UnfulfillablePreviewItem](docs/UnfulfillablePreviewItem.md)
 - [UpdateFulfillmentOrderItem](docs/UpdateFulfillmentOrderItem.md)
 - [UpdateFulfillmentOrderRequest](docs/UpdateFulfillmentOrderRequest.md)
 - [UpdateFulfillmentOrderResponse](docs/UpdateFulfillmentOrderResponse.md)
 - [Weight](docs/Weight.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

