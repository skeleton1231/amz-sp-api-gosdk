# {{classname}}

All URIs are relative to *https://sellingpartnerapi-na.amazon.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReport**](ReportsApi.md#CancelReport) | **Delete** /reports/2020-09-04/reports/{reportId} | 
[**CancelReportSchedule**](ReportsApi.md#CancelReportSchedule) | **Delete** /reports/2020-09-04/schedules/{reportScheduleId} | 
[**CreateReport**](ReportsApi.md#CreateReport) | **Post** /reports/2020-09-04/reports | 
[**CreateReportSchedule**](ReportsApi.md#CreateReportSchedule) | **Post** /reports/2020-09-04/schedules | 
[**GetReport**](ReportsApi.md#GetReport) | **Get** /reports/2020-09-04/reports/{reportId} | 
[**GetReportDocument**](ReportsApi.md#GetReportDocument) | **Get** /reports/2020-09-04/documents/{reportDocumentId} | 
[**GetReportSchedule**](ReportsApi.md#GetReportSchedule) | **Get** /reports/2020-09-04/schedules/{reportScheduleId} | 
[**GetReportSchedules**](ReportsApi.md#GetReportSchedules) | **Get** /reports/2020-09-04/schedules | 
[**GetReports**](ReportsApi.md#GetReports) | **Get** /reports/2020-09-04/reports | 

# **CancelReport**
> CancelReportResponse CancelReport(ctx, reportId)


Effective **June 27, 2023**, the `cancelReport` operation will no longer be available in the Selling Partner API for Reports v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **string**| The identifier for the report. This identifier is unique only in combination with a seller ID. | 

### Return type

[**CancelReportResponse**](CancelReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelReportSchedule**
> CancelReportScheduleResponse CancelReportSchedule(ctx, reportScheduleId)


Effective **June 27, 2023**, the `cancelReportSchedule` operation will no longer be available in the Selling Partner API for Reports v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportScheduleId** | **string**| The identifier for the report schedule. This identifier is unique only in combination with a seller ID. | 

### Return type

[**CancelReportScheduleResponse**](CancelReportScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateReport**
> CreateReportResponse CreateReport(ctx, body)


Effective **June 27, 2023**, the `createReport` operation will no longer be available in the Selling Partner API for Reports v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateReportSpecification**](CreateReportSpecification.md)|  | 

### Return type

[**CreateReportResponse**](CreateReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateReportSchedule**
> CreateReportScheduleResponse CreateReportSchedule(ctx, body)


Effective **June 27, 2023**, the `createReportSchedule` operation will no longer be available in the Selling Partner API for Reports v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateReportScheduleSpecification**](CreateReportScheduleSpecification.md)|  | 

### Return type

[**CreateReportScheduleResponse**](CreateReportScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReport**
> GetReportResponse GetReport(ctx, reportId)


Effective **June 27, 2023**, the `getReport` operation will no longer be available in the Selling Partner API for Reports v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **string**| The identifier for the report. This identifier is unique only in combination with a seller ID. | 

### Return type

[**GetReportResponse**](GetReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportDocument**
> GetReportDocumentResponse GetReportDocument(ctx, reportDocumentId)


Effective **June 27, 2023**, the `getReportDocument` operation will no longer be available in the Selling Partner API for Reports v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportDocumentId** | **string**| The identifier for the report document. | 

### Return type

[**GetReportDocumentResponse**](GetReportDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportSchedule**
> GetReportScheduleResponse GetReportSchedule(ctx, reportScheduleId)


Effective **June 27, 2023**, the `getReportSchedule` operation will no longer be available in the Selling Partner API for Reports v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportScheduleId** | **string**| The identifier for the report schedule. This identifier is unique only in combination with a seller ID. | 

### Return type

[**GetReportScheduleResponse**](GetReportScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportSchedules**
> GetReportSchedulesResponse GetReportSchedules(ctx, reportTypes)


Effective **June 27, 2023**, the `getReportSchedules` operation will no longer be available in the Selling Partner API for Reports v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportTypes** | [**[]string**](string.md)| A list of report types used to filter report schedules. | 

### Return type

[**GetReportSchedulesResponse**](GetReportSchedulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReports**
> GetReportsResponse GetReports(ctx, optional)


Effective **June 27, 2023**, the `getReports` operation will no longer be available in the Selling Partner API for Reports v2020-09-04 and all calls to it will fail. Integrations that rely on this operation should migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiGetReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportTypes** | [**optional.Interface of []string**](string.md)| A list of report types used to filter reports. When reportTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either reportTypes or nextToken is required. | 
 **processingStatuses** | [**optional.Interface of []string**](string.md)| A list of processing statuses used to filter reports. | 
 **marketplaceIds** | [**optional.Interface of []string**](string.md)| A list of marketplace identifiers used to filter reports. The reports returned will match at least one of the marketplaces that you specify. | 
 **pageSize** | **optional.Int32**| The maximum number of reports to return in a single call. | [default to 10]
 **createdSince** | **optional.Time**| The earliest report creation date and time for reports to include in the response, in ISO 8601 date time format. The default is 90 days ago. Reports are retained for a maximum of 90 days. | 
 **createdUntil** | **optional.Time**| The latest report creation date and time for reports to include in the response, in ISO 8601 date time format. The default is now. | 
 **nextToken** | **optional.String**| A string token returned in the response to your previous request. nextToken is returned when the number of results exceeds the specified pageSize value. To get the next page of results, call the getReports operation and include this token as the only parameter. Specifying nextToken with any other parameters will cause the request to fail. | 

### Return type

[**GetReportsResponse**](GetReportsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

