# Go API client for swagger

Provides programmatic access to Amazon Shipping APIs.   **Note:** If you are new to the Amazon Shipping API, refer to the latest version of <a href=\"https://developer-docs.amazon.com/amazon-shipping/docs/shipping-api-v2-reference\">Amazon Shipping API (v2)</a> on the <a href=\"https://developer-docs.amazon.com/amazon-shipping/\">Amazon Shipping Developer Documentation</a> site.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1
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
*ShippingApi* | [**CancelShipment**](docs/ShippingApi.md#cancelshipment) | **Post** /shipping/v1/shipments/{shipmentId}/cancel | 
*ShippingApi* | [**CreateShipment**](docs/ShippingApi.md#createshipment) | **Post** /shipping/v1/shipments | 
*ShippingApi* | [**GetAccount**](docs/ShippingApi.md#getaccount) | **Get** /shipping/v1/account | 
*ShippingApi* | [**GetRates**](docs/ShippingApi.md#getrates) | **Post** /shipping/v1/rates | 
*ShippingApi* | [**GetShipment**](docs/ShippingApi.md#getshipment) | **Get** /shipping/v1/shipments/{shipmentId} | 
*ShippingApi* | [**GetTrackingInformation**](docs/ShippingApi.md#gettrackinginformation) | **Get** /shipping/v1/tracking/{trackingId} | 
*ShippingApi* | [**PurchaseLabels**](docs/ShippingApi.md#purchaselabels) | **Post** /shipping/v1/shipments/{shipmentId}/purchaseLabels | 
*ShippingApi* | [**PurchaseShipment**](docs/ShippingApi.md#purchaseshipment) | **Post** /shipping/v1/purchaseShipment | 
*ShippingApi* | [**RetrieveShippingLabel**](docs/ShippingApi.md#retrieveshippinglabel) | **Post** /shipping/v1/shipments/{shipmentId}/containers/{trackingId}/label | 

## Documentation For Models

 - [AcceptedRate](docs/AcceptedRate.md)
 - [Account](docs/Account.md)
 - [Address](docs/Address.md)
 - [CancelShipmentResponse](docs/CancelShipmentResponse.md)
 - [Container](docs/Container.md)
 - [ContainerItem](docs/ContainerItem.md)
 - [ContainerSpecification](docs/ContainerSpecification.md)
 - [CreateShipmentRequest](docs/CreateShipmentRequest.md)
 - [CreateShipmentResponse](docs/CreateShipmentResponse.md)
 - [CreateShipmentResult](docs/CreateShipmentResult.md)
 - [Currency](docs/Currency.md)
 - [Dimensions](docs/Dimensions.md)
 - [Event](docs/Event.md)
 - [GetAccountResponse](docs/GetAccountResponse.md)
 - [GetRatesRequest](docs/GetRatesRequest.md)
 - [GetRatesResponse](docs/GetRatesResponse.md)
 - [GetRatesResult](docs/GetRatesResult.md)
 - [GetShipmentResponse](docs/GetShipmentResponse.md)
 - [GetTrackingInformationResponse](docs/GetTrackingInformationResponse.md)
 - [Label](docs/Label.md)
 - [LabelResult](docs/LabelResult.md)
 - [LabelSpecification](docs/LabelSpecification.md)
 - [Location](docs/Location.md)
 - [ModelError](docs/ModelError.md)
 - [Party](docs/Party.md)
 - [PurchaseLabelsRequest](docs/PurchaseLabelsRequest.md)
 - [PurchaseLabelsResponse](docs/PurchaseLabelsResponse.md)
 - [PurchaseLabelsResult](docs/PurchaseLabelsResult.md)
 - [PurchaseShipmentRequest](docs/PurchaseShipmentRequest.md)
 - [PurchaseShipmentResponse](docs/PurchaseShipmentResponse.md)
 - [PurchaseShipmentResult](docs/PurchaseShipmentResult.md)
 - [Rate](docs/Rate.md)
 - [RetrieveShippingLabelRequest](docs/RetrieveShippingLabelRequest.md)
 - [RetrieveShippingLabelResponse](docs/RetrieveShippingLabelResponse.md)
 - [RetrieveShippingLabelResult](docs/RetrieveShippingLabelResult.md)
 - [ServiceRate](docs/ServiceRate.md)
 - [ServiceType](docs/ServiceType.md)
 - [Shipment](docs/Shipment.md)
 - [ShippingPromiseSet](docs/ShippingPromiseSet.md)
 - [TimeRange](docs/TimeRange.md)
 - [TrackingInformation](docs/TrackingInformation.md)
 - [TrackingSummary](docs/TrackingSummary.md)
 - [Weight](docs/Weight.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


