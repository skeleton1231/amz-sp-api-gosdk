# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAppointmentForServiceJobByServiceJobId**](ServiceApi.md#AddAppointmentForServiceJobByServiceJobId) | **Post** /service/v1/serviceJobs/{serviceJobId}/appointments | 
[**AssignAppointmentResources**](ServiceApi.md#AssignAppointmentResources) | **Put** /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}/resources | 
[**CancelReservation**](ServiceApi.md#CancelReservation) | **Delete** /service/v1/reservation/{reservationId} | 
[**CancelServiceJobByServiceJobId**](ServiceApi.md#CancelServiceJobByServiceJobId) | **Put** /service/v1/serviceJobs/{serviceJobId}/cancellations | 
[**CompleteServiceJobByServiceJobId**](ServiceApi.md#CompleteServiceJobByServiceJobId) | **Put** /service/v1/serviceJobs/{serviceJobId}/completions | 
[**CreateReservation**](ServiceApi.md#CreateReservation) | **Post** /service/v1/reservation | 
[**CreateServiceDocumentUploadDestination**](ServiceApi.md#CreateServiceDocumentUploadDestination) | **Post** /service/v1/documents | 
[**GetAppointmentSlots**](ServiceApi.md#GetAppointmentSlots) | **Get** /service/v1/appointmentSlots | 
[**GetAppointmmentSlotsByJobId**](ServiceApi.md#GetAppointmmentSlotsByJobId) | **Get** /service/v1/serviceJobs/{serviceJobId}/appointmentSlots | 
[**GetFixedSlotCapacity**](ServiceApi.md#GetFixedSlotCapacity) | **Post** /service/v1/serviceResources/{resourceId}/capacity/fixed | 
[**GetRangeSlotCapacity**](ServiceApi.md#GetRangeSlotCapacity) | **Post** /service/v1/serviceResources/{resourceId}/capacity/range | 
[**GetServiceJobByServiceJobId**](ServiceApi.md#GetServiceJobByServiceJobId) | **Get** /service/v1/serviceJobs/{serviceJobId} | 
[**GetServiceJobs**](ServiceApi.md#GetServiceJobs) | **Get** /service/v1/serviceJobs | 
[**RescheduleAppointmentForServiceJobByServiceJobId**](ServiceApi.md#RescheduleAppointmentForServiceJobByServiceJobId) | **Post** /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId} | 
[**SetAppointmentFulfillmentData**](ServiceApi.md#SetAppointmentFulfillmentData) | **Put** /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}/fulfillment | 
[**UpdateReservation**](ServiceApi.md#UpdateReservation) | **Put** /service/v1/reservation/{reservationId} | 
[**UpdateSchedule**](ServiceApi.md#UpdateSchedule) | **Put** /service/v1/serviceResources/{resourceId}/schedules | 

# **AddAppointmentForServiceJobByServiceJobId**
> SetAppointmentResponse AddAppointmentForServiceJobByServiceJobId(ctx, body, serviceJobId)


Adds an appointment to the service job indicated by the service job identifier specified.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddAppointmentRequest**](AddAppointmentRequest.md)| Add appointment operation input details. | 
  **serviceJobId** | **string**| An Amazon defined service job identifier. | 

### Return type

[**SetAppointmentResponse**](SetAppointmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignAppointmentResources**
> AssignAppointmentResourcesResponse AssignAppointmentResources(ctx, body, serviceJobId, appointmentId)


Assigns new resource(s) or overwrite/update the existing one(s) to a service job appointment.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 1 | 2 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AssignAppointmentResourcesRequest**](AssignAppointmentResourcesRequest.md)|  | 
  **serviceJobId** | **string**| An Amazon-defined service job identifier. Get this value by calling the &#x60;getServiceJobs&#x60; operation of the Services API. | 
  **appointmentId** | **string**| An Amazon-defined identifier of active service job appointment. | 

### Return type

[**AssignAppointmentResourcesResponse**](AssignAppointmentResourcesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelReservation**
> CancelReservationResponse CancelReservation(ctx, reservationId, marketplaceIds)


Cancel a reservation.   **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reservationId** | **string**| Reservation Identifier | 
  **marketplaceIds** | [**[]string**](string.md)| An identifier for the marketplace in which the resource operates. | 

### Return type

[**CancelReservationResponse**](CancelReservationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelServiceJobByServiceJobId**
> CancelServiceJobByServiceJobIdResponse CancelServiceJobByServiceJobId(ctx, serviceJobId, cancellationReasonCode)


Cancels the service job indicated by the service job identifier specified.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceJobId** | **string**| An Amazon defined service job identifier. | 
  **cancellationReasonCode** | **string**| A cancel reason code that specifies the reason for cancelling a service job. | 

### Return type

[**CancelServiceJobByServiceJobIdResponse**](CancelServiceJobByServiceJobIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompleteServiceJobByServiceJobId**
> CompleteServiceJobByServiceJobIdResponse CompleteServiceJobByServiceJobId(ctx, serviceJobId)


Completes the service job indicated by the service job identifier specified.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceJobId** | **string**| An Amazon defined service job identifier. | 

### Return type

[**CompleteServiceJobByServiceJobIdResponse**](CompleteServiceJobByServiceJobIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateReservation**
> CreateReservationResponse CreateReservation(ctx, body, marketplaceIds)


Create a reservation.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateReservationRequest**](CreateReservationRequest.md)| Reservation details | 
  **marketplaceIds** | [**[]string**](string.md)| An identifier for the marketplace in which the resource operates. | 

### Return type

[**CreateReservationResponse**](CreateReservationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServiceDocumentUploadDestination**
> CreateServiceDocumentUploadDestination CreateServiceDocumentUploadDestination(ctx, body)


Creates an upload destination.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceUploadDocument**](ServiceUploadDocument.md)| Upload document operation input details. | 

### Return type

[**CreateServiceDocumentUploadDestination**](CreateServiceDocumentUploadDestination.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppointmentSlots**
> GetAppointmentSlotsResponse GetAppointmentSlots(ctx, asin, storeId, marketplaceIds, optional)


Gets appointment slots as per the service context specified.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 20 | 40 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asin** | **string**| ASIN associated with the service. | 
  **storeId** | **string**| Store identifier defining the region scope to retrive appointment slots. | 
  **marketplaceIds** | [**[]string**](string.md)| An identifier for the marketplace for which appointment slots are queried | 
 **optional** | ***ServiceApiGetAppointmentSlotsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetAppointmentSlotsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startTime** | **optional.String**| A time from which the appointment slots will be retrieved. The specified time must be in ISO 8601 format. If &#x60;startTime&#x60; is provided, &#x60;endTime&#x60; should also be provided. Default value is as per business configuration. | 
 **endTime** | **optional.String**| A time up to which the appointment slots will be retrieved. The specified time must be in ISO 8601 format. If &#x60;endTime&#x60; is provided, &#x60;startTime&#x60; should also be provided. Default value is as per business configuration. Maximum range of appointment slots can be 90 days. | 

### Return type

[**GetAppointmentSlotsResponse**](GetAppointmentSlotsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppointmmentSlotsByJobId**
> GetAppointmentSlotsResponse GetAppointmmentSlotsByJobId(ctx, serviceJobId, marketplaceIds, optional)


Gets appointment slots for the service associated with the service job id specified.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceJobId** | **string**| A service job identifier to retrive appointment slots for associated service. | 
  **marketplaceIds** | [**[]string**](string.md)| An identifier for the marketplace in which the resource operates. | 
 **optional** | ***ServiceApiGetAppointmmentSlotsByJobIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetAppointmmentSlotsByJobIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **optional.String**| A time from which the appointment slots will be retrieved. The specified time must be in ISO 8601 format. If &#x60;startTime&#x60; is provided, &#x60;endTime&#x60; should also be provided. Default value is as per business configuration. | 
 **endTime** | **optional.String**| A time up to which the appointment slots will be retrieved. The specified time must be in ISO 8601 format. If &#x60;endTime&#x60; is provided, &#x60;startTime&#x60; should also be provided. Default value is as per business configuration. Maximum range of appointment slots can be 90 days. | 

### Return type

[**GetAppointmentSlotsResponse**](GetAppointmentSlotsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFixedSlotCapacity**
> FixedSlotCapacity GetFixedSlotCapacity(ctx, body, resourceId, marketplaceIds, optional)


Provides capacity in fixed-size slots.   **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FixedSlotCapacityQuery**](FixedSlotCapacityQuery.md)| Request body. | 
  **resourceId** | **string**| Resource Identifier. | 
  **marketplaceIds** | [**[]string**](string.md)| An identifier for the marketplace in which the resource operates. | 
 **optional** | ***ServiceApiGetFixedSlotCapacityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetFixedSlotCapacityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nextPageToken** | **optional.**| Next page token returned in the response of your previous request. | 

### Return type

[**FixedSlotCapacity**](FixedSlotCapacity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRangeSlotCapacity**
> RangeSlotCapacity GetRangeSlotCapacity(ctx, body, resourceId, marketplaceIds, optional)


Provides capacity slots in a format similar to availability records.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RangeSlotCapacityQuery**](RangeSlotCapacityQuery.md)| Request body. | 
  **resourceId** | **string**| Resource Identifier. | 
  **marketplaceIds** | [**[]string**](string.md)| An identifier for the marketplace in which the resource operates. | 
 **optional** | ***ServiceApiGetRangeSlotCapacityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetRangeSlotCapacityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nextPageToken** | **optional.**| Next page token returned in the response of your previous request. | 

### Return type

[**RangeSlotCapacity**](RangeSlotCapacity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceJobByServiceJobId**
> GetServiceJobByServiceJobIdResponse GetServiceJobByServiceJobId(ctx, serviceJobId)


Gets details of service job indicated by the provided `serviceJobID`.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 20 | 40 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceJobId** | **string**| A service job identifier. | 

### Return type

[**GetServiceJobByServiceJobIdResponse**](GetServiceJobByServiceJobIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceJobs**
> GetServiceJobsResponse GetServiceJobs(ctx, marketplaceIds, optional)


Gets service job details for the specified filter query.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 40 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketplaceIds** | [**[]string**](string.md)| Used to select jobs that were placed in the specified marketplaces. | 
 **optional** | ***ServiceApiGetServiceJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceOrderIds** | [**optional.Interface of []string**](string.md)| List of service order ids for the query you want to perform.Max values supported 20. | 
 **serviceJobStatus** | [**optional.Interface of []string**](string.md)| A list of one or more job status by which to filter the list of jobs. | 
 **pageToken** | **optional.String**| String returned in the response of your previous request. | 
 **pageSize** | **optional.Int32**| A non-negative integer that indicates the maximum number of jobs to return in the list, Value must be 1 - 20. Default 20. | [default to 20]
 **sortField** | **optional.String**| Sort fields on which you want to sort the output. | 
 **sortOrder** | **optional.String**| Sort order for the query you want to perform. | 
 **createdAfter** | **optional.String**| A date used for selecting jobs created at or after a specified time. Must be in ISO 8601 format. Required if &#x60;LastUpdatedAfter&#x60; is not specified. Specifying both &#x60;CreatedAfter&#x60; and &#x60;LastUpdatedAfter&#x60; returns an error. | 
 **createdBefore** | **optional.String**| A date used for selecting jobs created at or before a specified time. Must be in ISO 8601 format. | 
 **lastUpdatedAfter** | **optional.String**| A date used for selecting jobs updated at or after a specified time. Must be in ISO 8601 format. Required if &#x60;createdAfter&#x60; is not specified. Specifying both &#x60;CreatedAfter&#x60; and &#x60;LastUpdatedAfter&#x60; returns an error. | 
 **lastUpdatedBefore** | **optional.String**| A date used for selecting jobs updated at or before a specified time. Must be in ISO 8601 format. | 
 **scheduleStartDate** | **optional.String**| A date used for filtering jobs schedules at or after a specified time. Must be in ISO 8601 format. Schedule end date should not be earlier than schedule start date. | 
 **scheduleEndDate** | **optional.String**| A date used for filtering jobs schedules at or before a specified time. Must be in ISO 8601 format. Schedule end date should not be earlier than schedule start date. | 
 **asins** | [**optional.Interface of []string**](string.md)| List of Amazon Standard Identification Numbers (ASIN) of the items. Max values supported is 20. | 
 **requiredSkills** | [**optional.Interface of []string**](string.md)| A defined set of related knowledge, skills, experience, tools, materials, and work processes common to service delivery for a set of products and/or service scenarios. Max values supported is 20. | 
 **storeIds** | [**optional.Interface of []string**](string.md)| List of Amazon-defined identifiers for the region scope. Max values supported is 50. | 

### Return type

[**GetServiceJobsResponse**](GetServiceJobsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RescheduleAppointmentForServiceJobByServiceJobId**
> SetAppointmentResponse RescheduleAppointmentForServiceJobByServiceJobId(ctx, body, serviceJobId, appointmentId)


Reschedules an appointment for the service job indicated by the service job identifier specified.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RescheduleAppointmentRequest**](RescheduleAppointmentRequest.md)| Reschedule appointment operation input details. | 
  **serviceJobId** | **string**| An Amazon defined service job identifier. | 
  **appointmentId** | **string**| An existing appointment identifier for the Service Job. | 

### Return type

[**SetAppointmentResponse**](SetAppointmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAppointmentFulfillmentData**
> string SetAppointmentFulfillmentData(ctx, body, serviceJobId, appointmentId)


Updates the appointment fulfillment data related to a given `jobID` and `appointmentID`.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SetAppointmentFulfillmentDataRequest**](SetAppointmentFulfillmentDataRequest.md)| Appointment fulfillment data collection details. | 
  **serviceJobId** | **string**| An Amazon-defined service job identifier. Get this value by calling the &#x60;getServiceJobs&#x60; operation of the Services API. | 
  **appointmentId** | **string**| An Amazon-defined identifier of active service job appointment. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReservation**
> UpdateReservationResponse UpdateReservation(ctx, body, reservationId, marketplaceIds)


Update a reservation.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateReservationRequest**](UpdateReservationRequest.md)| Reservation details | 
  **reservationId** | **string**| Reservation Identifier | 
  **marketplaceIds** | [**[]string**](string.md)| An identifier for the marketplace in which the resource operates. | 

### Return type

[**UpdateReservationResponse**](UpdateReservationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchedule**
> UpdateScheduleResponse UpdateSchedule(ctx, body, resourceId, marketplaceIds)


Update the schedule of the given resource.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 5 | 20 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateScheduleRequest**](UpdateScheduleRequest.md)| Schedule details | 
  **resourceId** | **string**| Resource (store) Identifier | 
  **marketplaceIds** | [**[]string**](string.md)| An identifier for the marketplace in which the resource operates. | 

### Return type

[**UpdateScheduleResponse**](UpdateScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

