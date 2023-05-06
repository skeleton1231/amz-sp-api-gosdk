# ReportDocument

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportDocumentId** | **string** | The identifier for the report document. This identifier is unique only in combination with a seller ID. | [default to null]
**Url** | **string** | A presigned URL for the report document. If &#x60;compressionAlgorithm&#x60; is not returned, you can download the report directly from this URL. This URL expires after 5 minutes. | [default to null]
**EncryptionDetails** | [***ReportDocumentEncryptionDetails**](ReportDocumentEncryptionDetails.md) |  | [default to null]
**CompressionAlgorithm** | **string** | If the report document contents have been compressed, the compression algorithm used is returned in this property and you must decompress the report when you download. Otherwise, you can download the report directly. Refer to [Step 2. Download and decrypt the report](doc:reports-api-v2020-09-04-use-case-guide#step-2-download-and-decrypt-the-report) in the use case guide, where sample code is provided. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

