/*
 * Selling Partner API for Reports
 *
 * Effective **June 27, 2023**, the Selling Partner API for Reports v2020-09-04 will no longer be available and all calls to it will fail. Integrations that rely on the Reports API must migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.
 *
 * API version: 2020-09-04
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ReportDocument struct {
	// The identifier for the report document. This identifier is unique only in combination with a seller ID.
	ReportDocumentId string `json:"reportDocumentId"`
	// A presigned URL for the report document. If `compressionAlgorithm` is not returned, you can download the report directly from this URL. This URL expires after 5 minutes.
	Url string `json:"url"`
	EncryptionDetails *ReportDocumentEncryptionDetails `json:"encryptionDetails"`
	// If the report document contents have been compressed, the compression algorithm used is returned in this property and you must decompress the report when you download. Otherwise, you can download the report directly. Refer to [Step 2. Download and decrypt the report](doc:reports-api-v2020-09-04-use-case-guide#step-2-download-and-decrypt-the-report) in the use case guide, where sample code is provided.
	CompressionAlgorithm string `json:"compressionAlgorithm,omitempty"`
}
