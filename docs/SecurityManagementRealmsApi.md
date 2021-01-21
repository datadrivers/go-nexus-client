# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveRealms**](SecurityManagementRealmsApi.md#GetActiveRealms) | **Get** /v1/security/realms/active | List the active realm IDs in order
[**GetRealms**](SecurityManagementRealmsApi.md#GetRealms) | **Get** /v1/security/realms/available | List the available realms
[**SetActiveRealms**](SecurityManagementRealmsApi.md#SetActiveRealms) | **Put** /v1/security/realms/active | Set the active security realms in the order they should be used

# **GetActiveRealms**
> []string GetActiveRealms(ctx, )
List the active realm IDs in order

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRealms**
> []RealmApiXo GetRealms(ctx, )
List the available realms

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]RealmApiXo**](RealmApiXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetActiveRealms**
> SetActiveRealms(ctx, optional)
Set the active security realms in the order they should be used

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementRealmsApiSetActiveRealmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementRealmsApiSetActiveRealmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []string**](string.md)| The realm IDs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

