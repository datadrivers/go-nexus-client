# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeOrder**](SecurityManagementLDAPApi.md#ChangeOrder) | **Post** /v1/security/ldap/change-order | Change LDAP server order
[**CreateLdapServer**](SecurityManagementLDAPApi.md#CreateLdapServer) | **Post** /v1/security/ldap | Create LDAP server
[**DeleteLdapServer**](SecurityManagementLDAPApi.md#DeleteLdapServer) | **Delete** /v1/security/ldap/{name} | Delete LDAP server
[**GetLdapServer**](SecurityManagementLDAPApi.md#GetLdapServer) | **Get** /v1/security/ldap/{name} | Get LDAP server
[**GetLdapServers**](SecurityManagementLDAPApi.md#GetLdapServers) | **Get** /v1/security/ldap | List LDAP servers
[**UpdateLdapServer**](SecurityManagementLDAPApi.md#UpdateLdapServer) | **Put** /v1/security/ldap/{name} | Update LDAP server

# **ChangeOrder**
> ChangeOrder(ctx, optional)
Change LDAP server order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementLDAPApiChangeOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementLDAPApiChangeOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []string**](string.md)| Ordered list of LDAP server names | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLdapServer**
> CreateLdapServer(ctx, optional)
Create LDAP server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityManagementLDAPApiCreateLdapServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementLDAPApiCreateLdapServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateLdapServerXo**](CreateLdapServerXo.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLdapServer**
> DeleteLdapServer(ctx, name)
Delete LDAP server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the LDAP server to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLdapServer**
> GetLdapServer(ctx, name)
Get LDAP server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the LDAP server to retrieve | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLdapServers**
> GetLdapServers(ctx, )
List LDAP servers

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

# **UpdateLdapServer**
> UpdateLdapServer(ctx, name, optional)
Update LDAP server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the LDAP server to update | 
 **optional** | ***SecurityManagementLDAPApiUpdateLdapServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityManagementLDAPApiUpdateLdapServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateLdapServerXo**](UpdateLdapServerXo.md)| Updated values of LDAP server | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

