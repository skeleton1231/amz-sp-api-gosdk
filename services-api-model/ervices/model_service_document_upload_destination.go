/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Information about an upload destination.
type ServiceDocumentUploadDestination struct {
	// The unique identifier to be used by APIs that reference the upload destination.
	UploadDestinationId string `json:"uploadDestinationId"`
	// The URL to which to upload the file.
	Url string `json:"url"`
	EncryptionDetails *EncryptionDetails `json:"encryptionDetails"`
	// The headers to include in the upload request.
	Headers *interface{} `json:"headers,omitempty"`
}
