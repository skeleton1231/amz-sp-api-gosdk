# Feed

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedId** | **string** | The identifier for the feed. This identifier is unique only in combination with a seller ID. | [default to null]
**FeedType** | **string** | The feed type. | [default to null]
**MarketplaceIds** | **[]string** | A list of identifiers for the marketplaces that the feed is applied to. | [optional] [default to null]
**CreatedTime** | [**time.Time**](time.Time.md) | The date and time when the feed was created, in ISO 8601 date time format. | [default to null]
**ProcessingStatus** | **string** | The processing status of the feed. | [default to null]
**ProcessingStartTime** | [**time.Time**](time.Time.md) | The date and time when feed processing started, in ISO 8601 date time format. | [optional] [default to null]
**ProcessingEndTime** | [**time.Time**](time.Time.md) | The date and time when feed processing completed, in ISO 8601 date time format. | [optional] [default to null]
**ResultFeedDocumentId** | **string** | The identifier for the feed document. This identifier is unique only in combination with a seller ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

