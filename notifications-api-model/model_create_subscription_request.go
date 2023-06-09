/*
 * Selling Partner API for Notifications
 *
 * The Selling Partner API for Notifications lets you subscribe to notifications that are relevant to a selling partner's business. Using this API you can create a destination to receive notifications, subscribe to notifications, delete notification subscriptions, and more.  For more information, see the [Notifications Use Case Guide](doc:notifications-api-v1-use-case-guide).
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The request schema for the createSubscription operation.
type CreateSubscriptionRequest struct {
	// The version of the payload object to be used in the notification.
	PayloadVersion string `json:"payloadVersion,omitempty"`
	// The identifier for the destination where notifications will be delivered.
	DestinationId string `json:"destinationId,omitempty"`
	ProcessingDirective *ProcessingDirective `json:"processingDirective,omitempty"`
}
