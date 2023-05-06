/*
 * Selling Partner API for Reports
 *
 * Effective **June 27, 2023**, the Selling Partner API for Reports v2020-09-04 will no longer be available and all calls to it will fail. Integrations that rely on the Reports API must migrate to [Reports v2021-06-30](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-reference) to avoid service disruption.
 *
 * API version: 2020-09-04
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The response for the getReport operation.
type GetReportResponse struct {
	Payload *Report `json:"payload,omitempty"`
	Errors *[]ModelError `json:"errors,omitempty"`
}
