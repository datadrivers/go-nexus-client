# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search**](SearchApi.md#Search) | **Get** /v1/search | Search components
[**SearchAndDownloadAssets**](SearchApi.md#SearchAndDownloadAssets) | **Get** /v1/search/assets/download | Search and download asset
[**SearchAssets**](SearchApi.md#SearchAssets) | **Get** /v1/search/assets | Search assets

# **Search**
> PageComponentXo Search(ctx, optional)
Search components

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continuationToken** | **optional.String**| A token returned by a prior request. If present, the next page of results are returned | 
 **sort** | **optional.String**| The field to sort the results against, if left empty, a sort based on match weight will be used. | 
 **direction** | **optional.String**| The direction to sort records in, defaults to ascending (&#x27;asc&#x27;) for all sort fields, except version, which defaults to descending (&#x27;desc&#x27;) | 
 **timeout** | **optional.Int32**| How long to wait for search results in seconds. If this value is not provided, the system default timeout will be used. | 
 **q** | **optional.String**| Query by keyword | 
 **repository** | **optional.String**| Repository name | 
 **format** | **optional.String**| Query by format | 
 **group** | **optional.String**| Component group | 
 **name** | **optional.String**| Component name | 
 **version** | **optional.String**| Component version | 
 **md5** | **optional.String**| Specific MD5 hash of component&#x27;s asset | 
 **sha1** | **optional.String**| Specific SHA-1 hash of component&#x27;s asset | 
 **sha256** | **optional.String**| Specific SHA-256 hash of component&#x27;s asset | 
 **sha512** | **optional.String**| Specific SHA-512 hash of component&#x27;s asset | 
 **prerelease** | **optional.String**| Prerelease version flag | 
 **conanBaseVersion** | **optional.String**| baseVersion | 
 **conanChannel** | **optional.String**| channel | 
 **dockerImageName** | **optional.String**| Docker image name | 
 **dockerImageTag** | **optional.String**| Docker image tag | 
 **dockerLayerId** | **optional.String**| Docker layer ID | 
 **dockerContentDigest** | **optional.String**| Docker content digest | 
 **mavenGroupId** | **optional.String**| Maven groupId | 
 **mavenArtifactId** | **optional.String**| Maven artifactId | 
 **mavenBaseVersion** | **optional.String**| Maven base version | 
 **mavenExtension** | **optional.String**| Maven extension of component&#x27;s asset | 
 **mavenClassifier** | **optional.String**| Maven classifier of component&#x27;s asset | 
 **npmScope** | **optional.String**| npm scope | 
 **nugetId** | **optional.String**| NuGet id | 
 **nugetTags** | **optional.String**| NuGet tags | 
 **p2PluginName** | **optional.String**| p2 plugin name | 
 **pypiClassifiers** | **optional.String**| PyPI classifiers | 
 **pypiDescription** | **optional.String**| PyPI description | 
 **pypiKeywords** | **optional.String**| PyPI keywords | 
 **pypiSummary** | **optional.String**| PyPI summary | 
 **rubygemsDescription** | **optional.String**| RubyGems description | 
 **rubygemsPlatform** | **optional.String**| RubyGems platform | 
 **rubygemsSummary** | **optional.String**| RubyGems summary | 
 **yumArchitecture** | **optional.String**| Yum architecture | 

### Return type

[**PageComponentXo**](PageComponentXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchAndDownloadAssets**
> SearchAndDownloadAssets(ctx, optional)
Search and download asset

Returns a 302 Found with location header field set to download URL. Unless a sort parameter is supplied, the search must return a single asset to receive download URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiSearchAndDownloadAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchAndDownloadAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| The field to sort the results against, if left empty and more than 1 result is returned, the request will fail. | 
 **direction** | **optional.String**| The direction to sort records in, defaults to ascending (&#x27;asc&#x27;) for all sort fields, except version, which defaults to descending (&#x27;desc&#x27;) | 
 **timeout** | **optional.Int32**| How long to wait for search results in seconds. If this value is not provided, the system default timeout will be used. | 
 **q** | **optional.String**| Query by keyword | 
 **repository** | **optional.String**| Repository name | 
 **format** | **optional.String**| Query by format | 
 **group** | **optional.String**| Component group | 
 **name** | **optional.String**| Component name | 
 **version** | **optional.String**| Component version | 
 **md5** | **optional.String**| Specific MD5 hash of component&#x27;s asset | 
 **sha1** | **optional.String**| Specific SHA-1 hash of component&#x27;s asset | 
 **sha256** | **optional.String**| Specific SHA-256 hash of component&#x27;s asset | 
 **sha512** | **optional.String**| Specific SHA-512 hash of component&#x27;s asset | 
 **prerelease** | **optional.String**| Prerelease version flag | 
 **conanBaseVersion** | **optional.String**| baseVersion | 
 **conanChannel** | **optional.String**| channel | 
 **dockerImageName** | **optional.String**| Docker image name | 
 **dockerImageTag** | **optional.String**| Docker image tag | 
 **dockerLayerId** | **optional.String**| Docker layer ID | 
 **dockerContentDigest** | **optional.String**| Docker content digest | 
 **mavenGroupId** | **optional.String**| Maven groupId | 
 **mavenArtifactId** | **optional.String**| Maven artifactId | 
 **mavenBaseVersion** | **optional.String**| Maven base version | 
 **mavenExtension** | **optional.String**| Maven extension of component&#x27;s asset | 
 **mavenClassifier** | **optional.String**| Maven classifier of component&#x27;s asset | 
 **npmScope** | **optional.String**| npm scope | 
 **nugetId** | **optional.String**| NuGet id | 
 **nugetTags** | **optional.String**| NuGet tags | 
 **p2PluginName** | **optional.String**| p2 plugin name | 
 **pypiClassifiers** | **optional.String**| PyPI classifiers | 
 **pypiDescription** | **optional.String**| PyPI description | 
 **pypiKeywords** | **optional.String**| PyPI keywords | 
 **pypiSummary** | **optional.String**| PyPI summary | 
 **rubygemsDescription** | **optional.String**| RubyGems description | 
 **rubygemsPlatform** | **optional.String**| RubyGems platform | 
 **rubygemsSummary** | **optional.String**| RubyGems summary | 
 **yumArchitecture** | **optional.String**| Yum architecture | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchAssets**
> PageAssetXo SearchAssets(ctx, optional)
Search assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiSearchAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continuationToken** | **optional.String**| A token returned by a prior request. If present, the next page of results are returned | 
 **sort** | **optional.String**| The field to sort the results against, if left empty, a sort based on match weight will be used. | 
 **direction** | **optional.String**| The direction to sort records in, defaults to ascending (&#x27;asc&#x27;) for all sort fields, except version, which defaults to descending (&#x27;desc&#x27;) | 
 **timeout** | **optional.Int32**| How long to wait for search results in seconds. If this value is not provided, the system default timeout will be used. | 
 **q** | **optional.String**| Query by keyword | 
 **repository** | **optional.String**| Repository name | 
 **format** | **optional.String**| Query by format | 
 **group** | **optional.String**| Component group | 
 **name** | **optional.String**| Component name | 
 **version** | **optional.String**| Component version | 
 **md5** | **optional.String**| Specific MD5 hash of component&#x27;s asset | 
 **sha1** | **optional.String**| Specific SHA-1 hash of component&#x27;s asset | 
 **sha256** | **optional.String**| Specific SHA-256 hash of component&#x27;s asset | 
 **sha512** | **optional.String**| Specific SHA-512 hash of component&#x27;s asset | 
 **prerelease** | **optional.String**| Prerelease version flag | 
 **conanBaseVersion** | **optional.String**| baseVersion | 
 **conanChannel** | **optional.String**| channel | 
 **dockerImageName** | **optional.String**| Docker image name | 
 **dockerImageTag** | **optional.String**| Docker image tag | 
 **dockerLayerId** | **optional.String**| Docker layer ID | 
 **dockerContentDigest** | **optional.String**| Docker content digest | 
 **mavenGroupId** | **optional.String**| Maven groupId | 
 **mavenArtifactId** | **optional.String**| Maven artifactId | 
 **mavenBaseVersion** | **optional.String**| Maven base version | 
 **mavenExtension** | **optional.String**| Maven extension of component&#x27;s asset | 
 **mavenClassifier** | **optional.String**| Maven classifier of component&#x27;s asset | 
 **npmScope** | **optional.String**| npm scope | 
 **nugetId** | **optional.String**| NuGet id | 
 **nugetTags** | **optional.String**| NuGet tags | 
 **p2PluginName** | **optional.String**| p2 plugin name | 
 **pypiClassifiers** | **optional.String**| PyPI classifiers | 
 **pypiDescription** | **optional.String**| PyPI description | 
 **pypiKeywords** | **optional.String**| PyPI keywords | 
 **pypiSummary** | **optional.String**| PyPI summary | 
 **rubygemsDescription** | **optional.String**| RubyGems description | 
 **rubygemsPlatform** | **optional.String**| RubyGems platform | 
 **rubygemsSummary** | **optional.String**| RubyGems summary | 
 **yumArchitecture** | **optional.String**| Yum architecture | 

### Return type

[**PageAssetXo**](PageAssetXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

