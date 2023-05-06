# Go API client for swagger

The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2021-12-28
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
*VendorOrdersApi* | [**GetOrder**](docs/VendorOrdersApi.md#getorder) | **Get** /vendor/directFulfillment/orders/2021-12-28/purchaseOrders/{purchaseOrderNumber} | 
*VendorOrdersApi* | [**GetOrders**](docs/VendorOrdersApi.md#getorders) | **Get** /vendor/directFulfillment/orders/2021-12-28/purchaseOrders | 
*VendorOrdersApi* | [**SubmitAcknowledgement**](docs/VendorOrdersApi.md#submitacknowledgement) | **Post** /vendor/directFulfillment/orders/2021-12-28/acknowledgements | 

## Documentation For Models

 - [AcknowledgementStatus](docs/AcknowledgementStatus.md)
 - [Address](docs/Address.md)
 - [BuyerCustomizedInfoDetail](docs/BuyerCustomizedInfoDetail.md)
 - [ErrorList](docs/ErrorList.md)
 - [GiftDetails](docs/GiftDetails.md)
 - [ItemQuantity](docs/ItemQuantity.md)
 - [ModelError](docs/ModelError.md)
 - [Money](docs/Money.md)
 - [Order](docs/Order.md)
 - [OrderAcknowledgementItem](docs/OrderAcknowledgementItem.md)
 - [OrderDetails](docs/OrderDetails.md)
 - [OrderItem](docs/OrderItem.md)
 - [OrderItemAcknowledgement](docs/OrderItemAcknowledgement.md)
 - [OrderList](docs/OrderList.md)
 - [Pagination](docs/Pagination.md)
 - [PartyIdentification](docs/PartyIdentification.md)
 - [ScheduledDeliveryShipment](docs/ScheduledDeliveryShipment.md)
 - [ShipmentDates](docs/ShipmentDates.md)
 - [ShipmentDetails](docs/ShipmentDetails.md)
 - [SubmitAcknowledgementRequest](docs/SubmitAcknowledgementRequest.md)
 - [SubmitAcknowledgementResponse](docs/SubmitAcknowledgementResponse.md)
 - [TaxDetails](docs/TaxDetails.md)
 - [TaxItemDetails](docs/TaxItemDetails.md)
 - [TaxRegistrationDetails](docs/TaxRegistrationDetails.md)
 - [TransactionId](docs/TransactionId.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

