# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoutingRule**](RoutingRulesApi.md#CreateRoutingRule) | **Post** /v1/routing-rules | Create a single routing rule
[**DeleteRoutingRule**](RoutingRulesApi.md#DeleteRoutingRule) | **Delete** /v1/routing-rules/{name} | Delete a single routing rule
[**GetRoutingRule**](RoutingRulesApi.md#GetRoutingRule) | **Get** /v1/routing-rules/{name} | Get a single routing rule
[**GetRoutingRules**](RoutingRulesApi.md#GetRoutingRules) | **Get** /v1/routing-rules | List routing rules
[**UpdateRoutingRule**](RoutingRulesApi.md#UpdateRoutingRule) | **Put** /v1/routing-rules/{name} | Update a single routing rule

# **CreateRoutingRule**
> CreateRoutingRule(ctx, body)
Create a single routing rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutingRuleXo**](RoutingRuleXo.md)| A routing rule configuration | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoutingRule**
> DeleteRoutingRule(ctx, name)
Delete a single routing rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the routing rule to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoutingRule**
> RoutingRuleXo GetRoutingRule(ctx, name)
Get a single routing rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the routing rule to get | 

### Return type

[**RoutingRuleXo**](RoutingRuleXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoutingRules**
> []RoutingRuleXo GetRoutingRules(ctx, )
List routing rules

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]RoutingRuleXo**](RoutingRuleXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoutingRule**
> UpdateRoutingRule(ctx, body, name)
Update a single routing rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutingRuleXo**](RoutingRuleXo.md)| A routing rule configuration | 
  **name** | **string**| The name of the routing rule to update | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

