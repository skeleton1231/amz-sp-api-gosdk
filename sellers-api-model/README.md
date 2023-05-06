# Go API client for swagger

The Selling Partner API for Sellers lets you retrieve information on behalf of sellers about their seller account, such as the marketplaces they participate in. Along with listing the marketplaces that a seller can sell in, the API also provides additional information about the marketplace such as the default language and the default currency. The API also provides seller-specific information such as whether the seller has suspended listings in that marketplace.

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
*SellersApi* | [**GetMarketplaceParticipations**](docs/SellersApi.md#getmarketplaceparticipations) | **Get** /sellers/v1/marketplaceParticipations | 

## Documentation For Models

 - [GetMarketplaceParticipationsResponse](docs/GetMarketplaceParticipationsResponse.md)
 - [Marketplace](docs/Marketplace.md)
 - [MarketplaceParticipation](docs/MarketplaceParticipation.md)
 - [ModelError](docs/ModelError.md)
 - [Participation](docs/Participation.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

