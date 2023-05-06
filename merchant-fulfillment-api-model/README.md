# Go API client for swagger

The Selling Partner API for Merchant Fulfillment helps you build applications that let sellers purchase shipping for non-Prime and Prime orders using Amazon’s Buy Shipping Services.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v0
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
*MerchantFulfillmentApi* | [**CancelShipment**](docs/MerchantFulfillmentApi.md#cancelshipment) | **Delete** /mfn/v0/shipments/{shipmentId} | 
*MerchantFulfillmentApi* | [**CancelShipmentOld**](docs/MerchantFulfillmentApi.md#cancelshipmentold) | **Put** /mfn/v0/shipments/{shipmentId}/cancel | 
*MerchantFulfillmentApi* | [**CreateShipment**](docs/MerchantFulfillmentApi.md#createshipment) | **Post** /mfn/v0/shipments | 
*MerchantFulfillmentApi* | [**GetAdditionalSellerInputs**](docs/MerchantFulfillmentApi.md#getadditionalsellerinputs) | **Post** /mfn/v0/additionalSellerInputs | 
*MerchantFulfillmentApi* | [**GetAdditionalSellerInputsOld**](docs/MerchantFulfillmentApi.md#getadditionalsellerinputsold) | **Post** /mfn/v0/sellerInputs | 
*MerchantFulfillmentApi* | [**GetEligibleShipmentServices**](docs/MerchantFulfillmentApi.md#geteligibleshipmentservices) | **Post** /mfn/v0/eligibleShippingServices | 
*MerchantFulfillmentApi* | [**GetEligibleShipmentServicesOld**](docs/MerchantFulfillmentApi.md#geteligibleshipmentservicesold) | **Post** /mfn/v0/eligibleServices | 
*MerchantFulfillmentApi* | [**GetShipment**](docs/MerchantFulfillmentApi.md#getshipment) | **Get** /mfn/v0/shipments/{shipmentId} | 

## Documentation For Models

 - [AdditionalInputs](docs/AdditionalInputs.md)
 - [AdditionalSellerInput](docs/AdditionalSellerInput.md)
 - [AdditionalSellerInputs](docs/AdditionalSellerInputs.md)
 - [Address](docs/Address.md)
 - [AvailableCarrierWillPickUpOption](docs/AvailableCarrierWillPickUpOption.md)
 - [AvailableDeliveryExperienceOption](docs/AvailableDeliveryExperienceOption.md)
 - [AvailableShippingServiceOptions](docs/AvailableShippingServiceOptions.md)
 - [CancelShipmentResponse](docs/CancelShipmentResponse.md)
 - [CarrierWillPickUpOption](docs/CarrierWillPickUpOption.md)
 - [Constraint](docs/Constraint.md)
 - [CreateShipmentRequest](docs/CreateShipmentRequest.md)
 - [CreateShipmentResponse](docs/CreateShipmentResponse.md)
 - [CurrencyAmount](docs/CurrencyAmount.md)
 - [DeliveryExperienceOption](docs/DeliveryExperienceOption.md)
 - [DeliveryExperienceType](docs/DeliveryExperienceType.md)
 - [FileContents](docs/FileContents.md)
 - [FileType](docs/FileType.md)
 - [GetAdditionalSellerInputsRequest](docs/GetAdditionalSellerInputsRequest.md)
 - [GetAdditionalSellerInputsResponse](docs/GetAdditionalSellerInputsResponse.md)
 - [GetAdditionalSellerInputsResult](docs/GetAdditionalSellerInputsResult.md)
 - [GetEligibleShipmentServicesRequest](docs/GetEligibleShipmentServicesRequest.md)
 - [GetEligibleShipmentServicesResponse](docs/GetEligibleShipmentServicesResponse.md)
 - [GetEligibleShipmentServicesResult](docs/GetEligibleShipmentServicesResult.md)
 - [GetShipmentResponse](docs/GetShipmentResponse.md)
 - [HazmatType](docs/HazmatType.md)
 - [InputTargetType](docs/InputTargetType.md)
 - [Item](docs/Item.md)
 - [ItemLevelFields](docs/ItemLevelFields.md)
 - [Label](docs/Label.md)
 - [LabelCustomization](docs/LabelCustomization.md)
 - [LabelDimensions](docs/LabelDimensions.md)
 - [LabelFormat](docs/LabelFormat.md)
 - [LabelFormatOption](docs/LabelFormatOption.md)
 - [LabelFormatOptionRequest](docs/LabelFormatOptionRequest.md)
 - [Length](docs/Length.md)
 - [ModelError](docs/ModelError.md)
 - [PackageDimensions](docs/PackageDimensions.md)
 - [PredefinedPackageDimensions](docs/PredefinedPackageDimensions.md)
 - [RejectedShippingService](docs/RejectedShippingService.md)
 - [SellerInputDefinition](docs/SellerInputDefinition.md)
 - [Shipment](docs/Shipment.md)
 - [ShipmentRequestDetails](docs/ShipmentRequestDetails.md)
 - [ShipmentStatus](docs/ShipmentStatus.md)
 - [ShippingOfferingFilter](docs/ShippingOfferingFilter.md)
 - [ShippingService](docs/ShippingService.md)
 - [ShippingServiceOptions](docs/ShippingServiceOptions.md)
 - [StandardIdForLabel](docs/StandardIdForLabel.md)
 - [TemporarilyUnavailableCarrier](docs/TemporarilyUnavailableCarrier.md)
 - [TermsAndConditionsNotAcceptedCarrier](docs/TermsAndConditionsNotAcceptedCarrier.md)
 - [UnitOfLength](docs/UnitOfLength.md)
 - [UnitOfWeight](docs/UnitOfWeight.md)
 - [Weight](docs/Weight.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

