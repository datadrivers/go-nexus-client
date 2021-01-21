# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContentSelector**](ContentSelectorsApi.md#CreateContentSelector) | **Post** /v1/security/content-selectors | Create a new content selector
[**DeleteContentSelector**](ContentSelectorsApi.md#DeleteContentSelector) | **Delete** /v1/security/content-selectors/{name} | Delete a content selector
[**GetContentSelector**](ContentSelectorsApi.md#GetContentSelector) | **Get** /v1/security/content-selectors/{name} | Get a content selector by name
[**GetContentSelectors**](ContentSelectorsApi.md#GetContentSelectors) | **Get** /v1/security/content-selectors | List content selectors
[**UpdateContentSelector**](ContentSelectorsApi.md#UpdateContentSelector) | **Put** /v1/security/content-selectors/{name} | Update a content selector

# **CreateContentSelector**
> CreateContentSelector(ctx, optional)
Create a new content selector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContentSelectorsApiCreateContentSelectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentSelectorsApiCreateContentSelectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ContentSelectorApiCreateRequest**](ContentSelectorApiCreateRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContentSelector**
> DeleteContentSelector(ctx, name)
Delete a content selector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentSelector**
> ContentSelectorApiResponse GetContentSelector(ctx, name)
Get a content selector by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The content selector name | 

### Return type

[**ContentSelectorApiResponse**](ContentSelectorApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentSelectors**
> []ContentSelectorApiResponse GetContentSelectors(ctx, )
List content selectors

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ContentSelectorApiResponse**](ContentSelectorApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContentSelector**
> UpdateContentSelector(ctx, name, optional)
Update a content selector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The content selector name | 
 **optional** | ***ContentSelectorsApiUpdateContentSelectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentSelectorsApiUpdateContentSelectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ContentSelectorApiUpdateRequest**](ContentSelectorApiUpdateRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

