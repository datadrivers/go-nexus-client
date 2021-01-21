# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePassword**](SecurityManagementUsersApi.md#ChangePassword) | **Put** /v1/security/users/{userId}/change-password | Change a user&#x27;s password.
[**CreateUser**](SecurityManagementUsersApi.md#CreateUser) | **Post** /v1/security/users | Create a new user in the default source.
[**DeleteUser**](SecurityManagementUsersApi.md#DeleteUser) | **Delete** /v1/security/users/{userId} | Delete a user.
[**GetUsers**](SecurityManagementUsersApi.md#GetUsers) | **Get** /v1/security/users | Retrieve a list of users. Note if the source is not &#x27;default&#x27; the response is limited to 100 users.
[**UpdateUser**](SecurityManagementUsersApi.md#UpdateUser) | **Put** /v1/security/users/{userId} | Update an existing user.

# **ChangePassword**
> ChangePassword(ctx, userId, optional)
Change a user's password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The userid the request should apply to. | 
 **optional** | ***SecurityManagementUsersApiChangePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementUsersApiChangePasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of string**](string.md)| The new password to use. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> ApiUser CreateUser(ctx, optional)
Create a new user in the default source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementUsersApiCreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementUsersApiCreateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiCreateUser**](ApiCreateUser.md)| A representation of the user to create. | 

### Return type

[**ApiUser**](ApiUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, userId)
Delete a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The userid the request should apply to. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> []ApiUser GetUsers(ctx, optional)
Retrieve a list of users. Note if the source is not 'default' the response is limited to 100 users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementUsersApiGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementUsersApiGetUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **optional.String**| An optional term to search userids for. | 
 **source** | **optional.String**| An optional user source to restrict the search to. | 

### Return type

[**[]ApiUser**](ApiUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UpdateUser(ctx, userId, optional)
Update an existing user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The userid the request should apply to. | 
 **optional** | ***SecurityManagementUsersApiUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementUsersApiUpdateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiUser**](ApiUser.md)| A representation of the user to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

