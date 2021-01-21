# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get1**](FormatsApi.md#Get1) | **Get** /v1/formats/{format}/upload-specs | Get upload field requirements for the desired format
[**Get2**](FormatsApi.md#Get2) | **Get** /v1/formats/upload-specs | Get upload field requirements for each supported format

# **Get1**
> UploadDefinitionXo Get1(ctx, format)
Get upload field requirements for the desired format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **format** | **string**| The desired repository format | 

### Return type

[**UploadDefinitionXo**](UploadDefinitionXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get2**
> []UploadDefinitionXo Get2(ctx, )
Get upload field requirements for each supported format

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UploadDefinitionXo**](UploadDefinitionXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

