# Report

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceIds** | **[]string** | A list of marketplace identifiers for the report. | [optional] [default to null]
**ReportId** | **string** | The identifier for the report. This identifier is unique only in combination with a seller ID. | [default to null]
**ReportType** | **string** | The report type. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. | [default to null]
**DataStartTime** | [**time.Time**](time.Time.md) | The start of a date and time range used for selecting the data to report. | [optional] [default to null]
**DataEndTime** | [**time.Time**](time.Time.md) | The end of a date and time range used for selecting the data to report. | [optional] [default to null]
**ReportScheduleId** | **string** | The identifier of the report schedule that created this report (if any). This identifier is unique only in combination with a seller ID. | [optional] [default to null]
**CreatedTime** | [**time.Time**](time.Time.md) | The date and time when the report was created. | [default to null]
**ProcessingStatus** | **string** | The processing status of the report. | [default to null]
**ProcessingStartTime** | [**time.Time**](time.Time.md) | The date and time when the report processing started, in ISO 8601 date time format. | [optional] [default to null]
**ProcessingEndTime** | [**time.Time**](time.Time.md) | The date and time when the report processing completed, in ISO 8601 date time format. | [optional] [default to null]
**ReportDocumentId** | **string** | The identifier for the report document. Pass this into the getReportDocument operation to get the information you will need to retrieve the report document&#x27;s contents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

