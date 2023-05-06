/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response schema for the `createServiceDocumentUploadDestination` operation.
type CreateServiceDocumentUploadDestination struct {
	Payload *ServiceDocumentUploadDestination `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}