# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Read**](SecurityManagementAnonymousAccessApi.md#Read) | **Get** /v1/security/anonymous | Get Anonymous Access settings
[**Update**](SecurityManagementAnonymousAccessApi.md#Update) | **Put** /v1/security/anonymous | Update Anonymous Access settings

# **Read**
> AnonymousAccessSettingsXo Read(ctx, )
Get Anonymous Access settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AnonymousAccessSettingsXo**](AnonymousAccessSettingsXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> AnonymousAccessSettingsXo Update(ctx, optional)
Update Anonymous Access settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementAnonymousAccessApiUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementAnonymousAccessApiUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AnonymousAccessSettingsXo**](AnonymousAccessSettingsXo.md)|  | 

### Return type

[**AnonymousAccessSettingsXo**](AnonymousAccessSettingsXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

