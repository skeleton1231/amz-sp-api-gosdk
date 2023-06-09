/*
 * Selling Partner API for Services
 *
 * With the Services API, you can build applications that help service providers get and modify their service orders and manage their resources.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Encryption details for required client-side encryption and decryption of document contents.
type EncryptionDetails struct {
	// The encryption standard required to encrypt or decrypt the document contents.
	Standard string `json:"standard"`
	// The vector to encrypt or decrypt the document contents using Cipher Block Chaining (CBC).
	InitializationVector string `json:"initializationVector"`
	// The encryption key used to encrypt or decrypt the document contents.
	Key string `json:"key"`
}
