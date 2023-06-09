# Go API client for swagger

The Selling Partner API for Product Type Definitions provides programmatic access to attribute and data requirements for product types in the Amazon catalog. Use this API to return the JSON Schema for a product type that you can then use with other Selling Partner APIs, such as the Selling Partner API for Listings Items, the Selling Partner API for Catalog Items, and the Selling Partner API for Feeds (for JSON-based listing feeds).  For more information, see the [Product Type Definitions API Use Case Guide](doc:product-type-api-use-case-guide).

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2020-09-01
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
*DefinitionsApi* | [**GetDefinitionsProductType**](docs/DefinitionsApi.md#getdefinitionsproducttype) | **Get** /definitions/2020-09-01/productTypes/{productType} | 
*DefinitionsApi* | [**SearchDefinitionsProductTypes**](docs/DefinitionsApi.md#searchdefinitionsproducttypes) | **Get** /definitions/2020-09-01/productTypes | 

## Documentation For Models

 - [ErrorList](docs/ErrorList.md)
 - [ModelError](docs/ModelError.md)
 - [ProductType](docs/ProductType.md)
 - [ProductTypeDefinition](docs/ProductTypeDefinition.md)
 - [ProductTypeList](docs/ProductTypeList.md)
 - [ProductTypeVersion](docs/ProductTypeVersion.md)
 - [PropertyGroup](docs/PropertyGroup.md)
 - [SchemaLink](docs/SchemaLink.md)
 - [SchemaLinkLink](docs/SchemaLinkLink.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


