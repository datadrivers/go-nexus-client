# HttpClientConnectionAttributes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retries** | **int32** | Total retries if the initial connection attempt suffers a timeout | [optional] [default to null]
**UserAgentSuffix** | **string** | Custom fragment to append to User-Agent header in HTTP requests | [optional] [default to null]
**Timeout** | **int32** | Seconds to wait for activity before stopping and retrying the connection | [optional] [default to null]
**EnableCircularRedirects** | **bool** | Whether to enable redirects to the same location (may be required by some servers) | [optional] [default to null]
**EnableCookies** | **bool** | Whether to allow cookies to be stored and used | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

