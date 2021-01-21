# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](SecurityManagementRolesApi.md#Create) | **Post** /v1/security/roles | Create role
[**Delete**](SecurityManagementRolesApi.md#Delete) | **Delete** /v1/security/roles/{id} | Delete role
[**GetRole**](SecurityManagementRolesApi.md#GetRole) | **Get** /v1/security/roles/{id} | Get role
[**GetRoles**](SecurityManagementRolesApi.md#GetRoles) | **Get** /v1/security/roles | List roles
[**Update1**](SecurityManagementRolesApi.md#Update1) | **Put** /v1/security/roles/{id} | Update role

# **Create**
> RoleXoResponse Create(ctx, body)
Create role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleXoRequest**](RoleXoRequest.md)| A role configuration | 

### Return type

[**RoleXoResponse**](RoleXOResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, id)
Delete role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the role to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> RoleXoResponse GetRole(ctx, id, optional)
Get role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the role to get | 
 **optional** | ***SecurityManagementRolesApiGetRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementRolesApiGetRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| The id of the user source to filter the roles by. Available sources can be fetched using the &#x27;User Sources&#x27; endpoint. | [default to default]

### Return type

[**RoleXoResponse**](RoleXOResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoles**
> []RoleXoResponse GetRoles(ctx, optional)
List roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementRolesApiGetRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementRolesApiGetRolesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **optional.String**| The id of the user source to filter the roles by, if supplied. Otherwise roles from all user sources will be returned. | 

### Return type

[**[]RoleXoResponse**](RoleXOResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update1**
> Update1(ctx, body, id)
Update role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleXoRequest**](RoleXoRequest.md)| A role configuration | 
  **id** | **string**| The id of the role to update | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

