# OrderMetricsInterval

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **string** | The interval of time based on requested granularity (ex. Hour, Day, etc.) If this is the first or the last interval from the list, it might contain incomplete data if the requested interval doesn&#x27;t align with the requested granularity (ex. request interval 2018-09-01T02:00:00Z--2018-09-04T19:00:00Z and granularity day will result in Sept 1st UTC day and Sept 4th UTC days having partial data). | [default to null]
**UnitCount** | **int32** | The number of units in orders based on the specified filters. | [default to null]
**OrderItemCount** | **int32** | The number of order items based on the specified filters. | [default to null]
**OrderCount** | **int32** | The number of orders based on the specified filters. | [default to null]
**AverageUnitPrice** | [***Money**](Money.md) |  | [default to null]
**TotalSales** | [***Money**](Money.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

