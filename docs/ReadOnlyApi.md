# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForceRelease**](ReadOnlyApi.md#ForceRelease) | **Post** /v1/read-only/force-release | Forcibly release read-only
[**Freeze**](ReadOnlyApi.md#Freeze) | **Post** /v1/read-only/freeze | Enable read-only
[**Get**](ReadOnlyApi.md#Get) | **Get** /v1/read-only | Get read-only state
[**Release**](ReadOnlyApi.md#Release) | **Post** /v1/read-only/release | Release read-only

# **ForceRelease**
> ForceRelease(ctx, )
Forcibly release read-only

Forcibly release read-only status, including System initiated tasks. Warning: may result in data loss.

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

# **Freeze**
> Freeze(ctx, )
Enable read-only

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

# **Get**
> ReadOnlyState Get(ctx, )
Get read-only state

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ReadOnlyState**](ReadOnlyState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Release**
> Release(ctx, )
Release read-only

Release administrator initiated read-only status. Will not release read-only caused by system tasks.

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

