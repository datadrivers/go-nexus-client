# S3BlobStoreApiBucket

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | The AWS region to create a new S3 bucket in or an existing S3 bucket&#x27;s region | [default to null]
**Name** | **string** | The name of the S3 bucket | [default to null]
**Prefix** | **string** | The S3 blob store (i.e S3 object) key prefix | [optional] [default to null]
**Expiration** | **int32** | How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable) | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

