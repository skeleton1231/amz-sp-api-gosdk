# CreateReportScheduleSpecification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | **string** | The report type. | [default to null]
**MarketplaceIds** | **[]string** | A list of marketplace identifiers for the report schedule. | [default to null]
**ReportOptions** | [***map[string]string**](map.md) |  | [optional] [default to null]
**Period** | **string** | One of a set of predefined ISO 8601 periods that specifies how often a report should be created. | [default to null]
**NextReportCreationTime** | [**time.Time**](time.Time.md) | The date and time when the schedule will create its next report, in ISO 8601 date time format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

