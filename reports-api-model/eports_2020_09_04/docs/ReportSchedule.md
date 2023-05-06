# ReportSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportScheduleId** | **string** | The identifier for the report schedule. This identifier is unique only in combination with a seller ID. | [default to null]
**ReportType** | **string** | The report type. | [default to null]
**MarketplaceIds** | **[]string** | A list of marketplace identifiers. The report document&#x27;s contents will contain data for all of the specified marketplaces, unless the report type indicates otherwise. | [optional] [default to null]
**ReportOptions** | [***map[string]string**](map.md) |  | [optional] [default to null]
**Period** | **string** | An ISO 8601 period value that indicates how often a report should be created. | [default to null]
**NextReportCreationTime** | [**time.Time**](time.Time.md) | The date and time when the schedule will create its next report, in ISO 8601 date time format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

