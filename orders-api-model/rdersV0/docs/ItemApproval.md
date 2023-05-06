# ItemApproval

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceId** | **int32** | Sequence number of the item approval. Each ItemApproval gets its sequenceId automatically from a monotonic increasing function. | [default to null]
**Timestamp** | **string** | Timestamp when the ItemApproval was recorded by Amazon&#x27;s internal order approvals system. In ISO 8601 date time format. | [default to null]
**Actor** | **string** | High level actors involved in the approval process. | [default to null]
**Approver** | **string** | Person or system that triggers the approval actions on behalf of the actor. | [optional] [default to null]
**ApprovalAction** | [***ItemApprovalAction**](ItemApprovalAction.md) |  | [default to null]
**ApprovalActionProcessStatus** | **string** | Status of approval action. | [default to null]
**ApprovalActionProcessStatusMessage** | **string** | Optional message to communicate optional additional context about the current status of the approval action. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

