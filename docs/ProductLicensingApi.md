# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicenseStatus**](ProductLicensingApi.md#GetLicenseStatus) | **Get** /v1/system/license | Get the current license status.
[**RemoveLicense**](ProductLicensingApi.md#RemoveLicense) | **Delete** /v1/system/license | Uninstall license if present.
[**SetLicense**](ProductLicensingApi.md#SetLicense) | **Post** /v1/system/license | Upload a new license file.

# **GetLicenseStatus**
> ApiLicenseDetailsXo GetLicenseStatus(ctx, )
Get the current license status.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiLicenseDetailsXo**](ApiLicenseDetailsXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveLicense**
> RemoveLicense(ctx, )
Uninstall license if present.

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

# **SetLicense**
> ApiLicenseDetailsXo SetLicense(ctx, optional)
Upload a new license file.

Server must be restarted to take effect

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductLicensingApiSetLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductLicensingApiSetLicenseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 

### Return type

[**ApiLicenseDetailsXo**](ApiLicenseDetailsXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

