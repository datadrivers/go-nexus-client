# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComponent**](ComponentsApi.md#DeleteComponent) | **Delete** /v1/components/{id} | Delete a single component
[**GetComponentById**](ComponentsApi.md#GetComponentById) | **Get** /v1/components/{id} | Get a single component
[**GetComponents**](ComponentsApi.md#GetComponents) | **Get** /v1/components | List components
[**UploadComponent**](ComponentsApi.md#UploadComponent) | **Post** /v1/components | Upload a single component

# **DeleteComponent**
> DeleteComponent(ctx, id)
Delete a single component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the component to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentById**
> ComponentXo GetComponentById(ctx, id)
Get a single component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the component to retrieve | 

### Return type

[**ComponentXo**](ComponentXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponents**
> PageComponentXo GetComponents(ctx, repository, optional)
List components

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| Repository from which you would like to retrieve components | 
 **optional** | ***ComponentsApiGetComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentsApiGetComponentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| A token returned by a prior request. If present, the next page of results are returned | 

### Return type

[**PageComponentXo**](PageComponentXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadComponent**
> UploadComponent(ctx, repository, optional)
Upload a single component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| Name of the repository to which you would like to upload the component | 
 **optional** | ***ComponentsApiUploadComponentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentsApiUploadComponentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rAsset** | **optional.Interface of *os.File****optional.**|  | 
 **rAssetPathId** | **optional.**|  | 
 **aptAsset** | **optional.Interface of *os.File****optional.**|  | 
 **yumDirectory** | **optional.**|  | 
 **yumAsset** | **optional.Interface of *os.File****optional.**|  | 
 **yumAssetFilename** | **optional.**|  | 
 **dockerAsset** | **optional.Interface of *os.File****optional.**|  | 
 **rubygemsAsset** | **optional.Interface of *os.File****optional.**|  | 
 **nugetAsset** | **optional.Interface of *os.File****optional.**|  | 
 **pypiAsset** | **optional.Interface of *os.File****optional.**|  | 
 **helmAsset** | **optional.Interface of *os.File****optional.**|  | 
 **npmAsset** | **optional.Interface of *os.File****optional.**|  | 
 **rawDirectory** | **optional.**|  | 
 **rawAsset1** | **optional.Interface of *os.File****optional.**|  | 
 **rawAsset1Filename** | **optional.**|  | 
 **rawAsset2** | **optional.Interface of *os.File****optional.**|  | 
 **rawAsset2Filename** | **optional.**|  | 
 **rawAsset3** | **optional.Interface of *os.File****optional.**|  | 
 **rawAsset3Filename** | **optional.**|  | 
 **maven2GroupId** | **optional.**|  | 
 **maven2ArtifactId** | **optional.**|  | 
 **maven2Version** | **optional.**|  | 
 **maven2GeneratePom** | **optional.**|  | 
 **maven2Packaging** | **optional.**|  | 
 **maven2Asset1** | **optional.Interface of *os.File****optional.**|  | 
 **maven2Asset1Classifier** | **optional.**|  | 
 **maven2Asset1Extension** | **optional.**|  | 
 **maven2Asset2** | **optional.Interface of *os.File****optional.**|  | 
 **maven2Asset2Classifier** | **optional.**|  | 
 **maven2Asset2Extension** | **optional.**|  | 
 **maven2Asset3** | **optional.Interface of *os.File****optional.**|  | 
 **maven2Asset3Classifier** | **optional.**|  | 
 **maven2Asset3Extension** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

