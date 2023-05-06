/*
 * Selling Partner API for Tokens 
 *
 * The Selling Partner API for Tokens provides a secure way to access a customer's PII (Personally Identifiable Information). You can call the Tokens API to get a Restricted Data Token (RDT) for one or more restricted resources that you specify. The RDT authorizes subsequent calls to restricted operations that correspond to the restricted resources that you specified.  For more information, see the [Tokens API Use Case Guide](doc:tokens-api-use-case-guide).
 *
 * API version: 2021-03-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The request schema for the createRestrictedDataToken operation.
type CreateRestrictedDataTokenRequest struct {
	// The application ID for the target application to which access is being delegated.
	TargetApplication string `json:"targetApplication,omitempty"`
	// A list of restricted resources. Maximum: 50
	RestrictedResources []RestrictedResource `json:"restrictedResources"`
}
