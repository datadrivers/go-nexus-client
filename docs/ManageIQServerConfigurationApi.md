# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableIq**](ManageIQServerConfigurationApi.md#DisableIq) | **Post** /v1/iq/disable | Disable IQ server
[**EnableIq**](ManageIQServerConfigurationApi.md#EnableIq) | **Post** /v1/iq/enable | Enable IQ server
[**GetConfiguration**](ManageIQServerConfigurationApi.md#GetConfiguration) | **Get** /v1/iq | Get IQ server configuration
[**UpdateConfiguration**](ManageIQServerConfigurationApi.md#UpdateConfiguration) | **Put** /v1/iq | Update IQ server configuration
[**VerifyConnection**](ManageIQServerConfigurationApi.md#VerifyConnection) | **Post** /v1/iq/verify-connection | Verify IQ server connection

# **DisableIq**
> DisableIq(ctx, )
Disable IQ server

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

# **EnableIq**
> EnableIq(ctx, )
Enable IQ server

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

# **GetConfiguration**
> GetConfiguration(ctx, )
Get IQ server configuration

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

# **UpdateConfiguration**
> UpdateConfiguration(ctx, optional)
Update IQ server configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManageIQServerConfigurationApiUpdateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManageIQServerConfigurationApiUpdateConfigurationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of IqConnectionXo**](IqConnectionXo.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyConnection**
> VerifyConnection(ctx, )
Verify IQ server connection

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

