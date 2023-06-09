# Go API client for swagger

The Selling Partner API for Retail Procurement Transaction Status provides programmatic access to status information on specific asynchronous POST transactions for vendors.

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
*VendorTransactionApi* | [**GetTransaction**](docs/VendorTransactionApi.md#gettransaction) | **Get** /vendor/transactions/v1/transactions/{transactionId} | 

## Documentation For Models

 - [GetTransactionResponse](docs/GetTransactionResponse.md)
 - [ModelError](docs/ModelError.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionStatus](docs/TransactionStatus.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


