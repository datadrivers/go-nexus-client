# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivilege**](SecurityManagementPrivilegesApi.md#CreatePrivilege) | **Post** /v1/security/privileges/wildcard | Create a wildcard type privilege.
[**CreatePrivilege1**](SecurityManagementPrivilegesApi.md#CreatePrivilege1) | **Post** /v1/security/privileges/application | Create an application type privilege.
[**CreatePrivilege2**](SecurityManagementPrivilegesApi.md#CreatePrivilege2) | **Post** /v1/security/privileges/repository-content-selector | Create a repository content selector type privilege.
[**CreatePrivilege3**](SecurityManagementPrivilegesApi.md#CreatePrivilege3) | **Post** /v1/security/privileges/repository-admin | Create a repository admin type privilege.
[**CreatePrivilege4**](SecurityManagementPrivilegesApi.md#CreatePrivilege4) | **Post** /v1/security/privileges/repository-view | Create a repository view type privilege.
[**CreatePrivilege5**](SecurityManagementPrivilegesApi.md#CreatePrivilege5) | **Post** /v1/security/privileges/script | Create a script type privilege.
[**DeletePrivilege**](SecurityManagementPrivilegesApi.md#DeletePrivilege) | **Delete** /v1/security/privileges/{privilegeId} | Delete a privilege by id.
[**GetPrivilege**](SecurityManagementPrivilegesApi.md#GetPrivilege) | **Get** /v1/security/privileges/{privilegeId} | Retrieve a privilege by id.
[**GetPrivileges**](SecurityManagementPrivilegesApi.md#GetPrivileges) | **Get** /v1/security/privileges | Retrieve a list of privileges.
[**UpdatePrivilege**](SecurityManagementPrivilegesApi.md#UpdatePrivilege) | **Put** /v1/security/privileges/wildcard/{privilegeId} | Update a wildcard type privilege.
[**UpdatePrivilege1**](SecurityManagementPrivilegesApi.md#UpdatePrivilege1) | **Put** /v1/security/privileges/application/{privilegeId} | Update an application type privilege.
[**UpdatePrivilege2**](SecurityManagementPrivilegesApi.md#UpdatePrivilege2) | **Put** /v1/security/privileges/repository-view/{privilegeId} | Update a repository view type privilege.
[**UpdatePrivilege3**](SecurityManagementPrivilegesApi.md#UpdatePrivilege3) | **Put** /v1/security/privileges/repository-content-selector/{privilegeId} | Update a repository content selector type privilege.
[**UpdatePrivilege4**](SecurityManagementPrivilegesApi.md#UpdatePrivilege4) | **Put** /v1/security/privileges/repository-admin/{privilegeId} | Update a repository admin type privilege.
[**UpdatePrivilege5**](SecurityManagementPrivilegesApi.md#UpdatePrivilege5) | **Put** /v1/security/privileges/script/{privilegeId} | Update a script type privilege.

# **CreatePrivilege**
> CreatePrivilege(ctx, optional)
Create a wildcard type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementPrivilegesApiCreatePrivilegeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiCreatePrivilegeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiPrivilegeWildcardRequest**](ApiPrivilegeWildcardRequest.md)| The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePrivilege1**
> CreatePrivilege1(ctx, optional)
Create an application type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementPrivilegesApiCreatePrivilege1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiCreatePrivilege1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiPrivilegeApplicationRequest**](ApiPrivilegeApplicationRequest.md)| The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePrivilege2**
> CreatePrivilege2(ctx, optional)
Create a repository content selector type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementPrivilegesApiCreatePrivilege2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiCreatePrivilege2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiPrivilegeRepositoryContentSelectorRequest**](ApiPrivilegeRepositoryContentSelectorRequest.md)| The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePrivilege3**
> CreatePrivilege3(ctx, optional)
Create a repository admin type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementPrivilegesApiCreatePrivilege3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiCreatePrivilege3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiPrivilegeRepositoryAdminRequest**](ApiPrivilegeRepositoryAdminRequest.md)| The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePrivilege4**
> CreatePrivilege4(ctx, optional)
Create a repository view type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementPrivilegesApiCreatePrivilege4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiCreatePrivilege4Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiPrivilegeRepositoryViewRequest**](ApiPrivilegeRepositoryViewRequest.md)| The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePrivilege5**
> CreatePrivilege5(ctx, optional)
Create a script type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementPrivilegesApiCreatePrivilege5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiCreatePrivilege5Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiPrivilegeScriptRequest**](ApiPrivilegeScriptRequest.md)| The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePrivilege**
> DeletePrivilege(ctx, privilegeId)
Delete a privilege by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **privilegeId** | **string**| The id of the privilege to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrivilege**
> ApiPrivilege GetPrivilege(ctx, privilegeId)
Retrieve a privilege by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **privilegeId** | **string**| The id of the privilege to retrieve. | 

### Return type

[**ApiPrivilege**](ApiPrivilege.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrivileges**
> []ApiPrivilege GetPrivileges(ctx, )
Retrieve a list of privileges.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiPrivilege**](ApiPrivilege.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrivilege**
> UpdatePrivilege(ctx, privilegeId, optional)
Update a wildcard type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **privilegeId** | **string**| The id of the privilege to update. | 
 **optional** | ***SecurityManagementPrivilegesApiUpdatePrivilegeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiUpdatePrivilegeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiPrivilegeWildcardRequest**](ApiPrivilegeWildcardRequest.md)| The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrivilege1**
> UpdatePrivilege1(ctx, privilegeId, optional)
Update an application type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **privilegeId** | **string**| The id of the privilege to update. | 
 **optional** | ***SecurityManagementPrivilegesApiUpdatePrivilege1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiUpdatePrivilege1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiPrivilegeApplicationRequest**](ApiPrivilegeApplicationRequest.md)| The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrivilege2**
> UpdatePrivilege2(ctx, privilegeId, optional)
Update a repository view type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **privilegeId** | **string**| The id of the privilege to update. | 
 **optional** | ***SecurityManagementPrivilegesApiUpdatePrivilege2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiUpdatePrivilege2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiPrivilegeRepositoryViewRequest**](ApiPrivilegeRepositoryViewRequest.md)| The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrivilege3**
> UpdatePrivilege3(ctx, privilegeId, optional)
Update a repository content selector type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **privilegeId** | **string**| The id of the privilege to update. | 
 **optional** | ***SecurityManagementPrivilegesApiUpdatePrivilege3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiUpdatePrivilege3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiPrivilegeRepositoryContentSelectorRequest**](ApiPrivilegeRepositoryContentSelectorRequest.md)| The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrivilege4**
> UpdatePrivilege4(ctx, privilegeId, optional)
Update a repository admin type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **privilegeId** | **string**| The id of the privilege to update. | 
 **optional** | ***SecurityManagementPrivilegesApiUpdatePrivilege4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiUpdatePrivilege4Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiPrivilegeRepositoryAdminRequest**](ApiPrivilegeRepositoryAdminRequest.md)| The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrivilege5**
> UpdatePrivilege5(ctx, privilegeId, optional)
Update a script type privilege.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **privilegeId** | **string**| The id of the privilege to update. | 
 **optional** | ***SecurityManagementPrivilegesApiUpdatePrivilege5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementPrivilegesApiUpdatePrivilege5Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiPrivilegeScriptRequest**](ApiPrivilegeScriptRequest.md)| The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

