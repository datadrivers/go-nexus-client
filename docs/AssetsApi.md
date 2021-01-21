# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAsset**](AssetsApi.md#DeleteAsset) | **Delete** /v1/assets/{id} | Delete a single asset
[**GetAssetById**](AssetsApi.md#GetAssetById) | **Get** /v1/assets/{id} | Get a single asset
[**GetAssets**](AssetsApi.md#GetAssets) | **Get** /v1/assets | List assets

# **DeleteAsset**
> DeleteAsset(ctx, id)
Delete a single asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the asset to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssetById**
> AssetXo GetAssetById(ctx, id)
Get a single asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the asset to get | 

### Return type

[**AssetXo**](AssetXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssets**
> PageAssetXo GetAssets(ctx, repository, optional)
List assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| Repository from which you would like to retrieve assets. | 
 **optional** | ***AssetsApiGetAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiGetAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| A token returned by a prior request. If present, the next page of results are returned | 

### Return type

[**PageAssetXo**](PageAssetXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

