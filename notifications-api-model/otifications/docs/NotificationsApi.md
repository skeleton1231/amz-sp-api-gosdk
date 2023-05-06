# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDestination**](NotificationsApi.md#CreateDestination) | **Post** /notifications/v1/destinations | 
[**CreateSubscription**](NotificationsApi.md#CreateSubscription) | **Post** /notifications/v1/subscriptions/{notificationType} | 
[**DeleteDestination**](NotificationsApi.md#DeleteDestination) | **Delete** /notifications/v1/destinations/{destinationId} | 
[**DeleteSubscriptionById**](NotificationsApi.md#DeleteSubscriptionById) | **Delete** /notifications/v1/subscriptions/{notificationType}/{subscriptionId} | 
[**GetDestination**](NotificationsApi.md#GetDestination) | **Get** /notifications/v1/destinations/{destinationId} | 
[**GetDestinations**](NotificationsApi.md#GetDestinations) | **Get** /notifications/v1/destinations | 
[**GetSubscription**](NotificationsApi.md#GetSubscription) | **Get** /notifications/v1/subscriptions/{notificationType} | 
[**GetSubscriptionById**](NotificationsApi.md#GetSubscriptionById) | **Get** /notifications/v1/subscriptions/{notificationType}/{subscriptionId} | 

# **CreateDestination**
> CreateDestinationResponse CreateDestination(ctx, body)


Creates a destination resource to receive notifications. The createDestination API is grantless. For more information, see [Grantless operations](doc:grantless-operations) in the Selling Partner API Developer Guide.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDestinationRequest**](CreateDestinationRequest.md)|  | 

### Return type

[**CreateDestinationResponse**](CreateDestinationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubscription**
> CreateSubscriptionResponse CreateSubscription(ctx, body, notificationType)


Creates a subscription for the specified notification type to be delivered to the specified destination. Before you can subscribe, you must first create the destination by calling the createDestination operation.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSubscriptionRequest**](CreateSubscriptionRequest.md)|  | 
  **notificationType** | **string**| The type of notification.   For more information about notification types, see [the Notifications API Use Case Guide](doc:notifications-api-v1-use-case-guide). | 

### Return type

[**CreateSubscriptionResponse**](CreateSubscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDestination**
> DeleteDestinationResponse DeleteDestination(ctx, destinationId)


Deletes the destination that you specify. The deleteDestination API is grantless. For more information, see [Grantless operations](doc:grantless-operations) in the Selling Partner API Developer Guide.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinationId** | **string**| The identifier for the destination that you want to delete. | 

### Return type

[**DeleteDestinationResponse**](DeleteDestinationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriptionById**
> DeleteSubscriptionByIdResponse DeleteSubscriptionById(ctx, subscriptionId, notificationType)


Deletes the subscription indicated by the subscription identifier and notification type that you specify. The subscription identifier can be for any subscription associated with your application. After you successfully call this operation, notifications will stop being sent for the associated subscription. The deleteSubscriptionById API is grantless. For more information, see [Grantless operations](doc:grantless-operations) in the Selling Partner API Developer Guide.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| The identifier for the subscription that you want to delete. | 
  **notificationType** | **string**| The type of notification.   For more information about notification types, see [the Notifications API Use Case Guide](doc:notifications-api-v1-use-case-guide). | 

### Return type

[**DeleteSubscriptionByIdResponse**](DeleteSubscriptionByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, Successful Operation Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDestination**
> GetDestinationResponse GetDestination(ctx, destinationId)


Returns information about the destination that you specify. The getDestination API is grantless. For more information, see [Grantless operations](doc:grantless-operations) in the Selling Partner API Developer Guide.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinationId** | **string**| The identifier generated when you created the destination. | 

### Return type

[**GetDestinationResponse**](GetDestinationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDestinations**
> GetDestinationsResponse GetDestinations(ctx, )


Returns information about all destinations. The getDestinations API is grantless. For more information, see [Grantless operations](doc:grantless-operations) in the Selling Partner API Developer Guide.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetDestinationsResponse**](GetDestinationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscription**
> GetSubscriptionResponse GetSubscription(ctx, notificationType)


Returns information about subscriptions of the specified notification type. You can use this API to get subscription information when you do not have a subscription identifier.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationType** | **string**| The type of notification.   For more information about notification types, see [the Notifications API Use Case Guide](doc:notifications-api-v1-use-case-guide). | 

### Return type

[**GetSubscriptionResponse**](GetSubscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptionById**
> GetSubscriptionByIdResponse GetSubscriptionById(ctx, subscriptionId, notificationType)


Returns information about a subscription for the specified notification type. The getSubscriptionById API is grantless. For more information, see [Grantless operations](doc:grantless-operations) in the Selling Partner API Developer Guide.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 5 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| The identifier for the subscription that you want to get. | 
  **notificationType** | **string**| The type of notification.   For more information about notification types, see [the Notifications API Use Case Guide](doc:notifications-api-v1-use-case-guide). | 

### Return type

[**GetSubscriptionByIdResponse**](GetSubscriptionByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

