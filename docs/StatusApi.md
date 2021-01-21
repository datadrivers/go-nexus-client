# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemStatusChecks**](StatusApi.md#GetSystemStatusChecks) | **Get** /v1/status/check | Health check endpoint that returns the results of the system status checks
[**IsAvailable**](StatusApi.md#IsAvailable) | **Get** /v1/status | Health check endpoint that validates server can respond to read requests
[**IsWritable**](StatusApi.md#IsWritable) | **Get** /v1/status/writable | Health check endpoint that validates server can respond to read and write requests

# **GetSystemStatusChecks**
> map[string]Result GetSystemStatusChecks(ctx, )
Health check endpoint that returns the results of the system status checks

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]Result**](Result.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsAvailable**
> IsAvailable(ctx, )
Health check endpoint that validates server can respond to read requests

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsWritable**
> IsWritable(ctx, )
Health check endpoint that validates server can respond to read and write requests

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

