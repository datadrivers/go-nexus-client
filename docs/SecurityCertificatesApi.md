# {{classname}}

All URIs are relative to */service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCertificate**](SecurityCertificatesApi.md#AddCertificate) | **Post** /v1/security/ssl/truststore | Add a certificate to the trust store.
[**GetTrustStoreCertificates**](SecurityCertificatesApi.md#GetTrustStoreCertificates) | **Get** /v1/security/ssl/truststore | Retrieve a list of certificates added to the trust store.
[**RemoveCertificate**](SecurityCertificatesApi.md#RemoveCertificate) | **Delete** /v1/security/ssl/truststore/{id} | Remove a certificate in the trust store.
[**RetrieveCertificate**](SecurityCertificatesApi.md#RetrieveCertificate) | **Get** /v1/security/ssl | Helper method to retrieve certificate details from a remote system.

# **AddCertificate**
> ApiCertificate AddCertificate(ctx, optional)
Add a certificate to the trust store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityCertificatesApiAddCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityCertificatesApiAddCertificateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of string**](string.md)| The certificate to add encoded in PEM format | 

### Return type

[**ApiCertificate**](ApiCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrustStoreCertificates**
> []ApiCertificate GetTrustStoreCertificates(ctx, )
Retrieve a list of certificates added to the trust store.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiCertificate**](ApiCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCertificate**
> RemoveCertificate(ctx, id)
Remove a certificate in the trust store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the certificate that should be removed. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCertificate**
> ApiCertificate RetrieveCertificate(ctx, host, optional)
Helper method to retrieve certificate details from a remote system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **host** | **string**| The remote system&#x27;s host name | 
 **optional** | ***SecurityCertificatesApiRetrieveCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityCertificatesApiRetrieveCertificateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **port** | **optional.Int32**| The port on the remote system to connect to | [default to 443]
 **protocolHint** | **optional.String**| An optional hint of the protocol to try for the connection | 

### Return type

[**ApiCertificate**](ApiCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

