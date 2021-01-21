# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEmailConfiguration**](EmailApi.md#DeleteEmailConfiguration) | **Delete** /v1/email | Disable and clear the email configuration
[**GetEmailConfiguration**](EmailApi.md#GetEmailConfiguration) | **Get** /v1/email | Retrieve the current email configuration
[**SetEmailConfiguration**](EmailApi.md#SetEmailConfiguration) | **Put** /v1/email | Set the current email configuration
[**TestEmailConfiguration**](EmailApi.md#TestEmailConfiguration) | **Post** /v1/email/verify | Send a test email to the email address provided in the request body

# **DeleteEmailConfiguration**
> DeleteEmailConfiguration(ctx, )
Disable and clear the email configuration

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

# **GetEmailConfiguration**
> ApiEmailConfiguration GetEmailConfiguration(ctx, )
Retrieve the current email configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEmailConfiguration**](ApiEmailConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetEmailConfiguration**
> SetEmailConfiguration(ctx, body)
Set the current email configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiEmailConfiguration**](ApiEmailConfiguration.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestEmailConfiguration**
> ApiEmailValidation TestEmailConfiguration(ctx, body)
Send a test email to the email address provided in the request body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| An email address to send a test email to | 

### Return type

[**ApiEmailValidation**](ApiEmailValidation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

