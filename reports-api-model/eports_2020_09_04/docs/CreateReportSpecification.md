# CreateReportSpecification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportOptions** | [***map[string]string**](map.md) |  | [optional] [default to null]
**ReportType** | **string** | The report type. | [default to null]
**DataStartTime** | [**time.Time**](time.Time.md) | The start of a date and time range, in ISO 8601 date time format, used for selecting the data to report. The default is now. The value must be prior to or equal to the current date and time. Not all report types make use of this. | [optional] [default to null]
**DataEndTime** | [**time.Time**](time.Time.md) | The end of a date and time range, in ISO 8601 date time format, used for selecting the data to report. The default is now. The value must be prior to or equal to the current date and time. Not all report types make use of this. | [optional] [default to null]
**MarketplaceIds** | **[]string** | A list of marketplace identifiers. The report document&#x27;s contents will contain data for all of the specified marketplaces, unless the report type indicates otherwise. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

