# \StorageApi

All URIs are relative to *https://api.growingbox.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStorageTemporaryCredentials**](StorageApi.md#GetStorageTemporaryCredentials) | **Get** /storage/v1/temporary_credentials | 获取上传文件临时凭证



## GetStorageTemporaryCredentials

> StorageTemporaryCredentials GetStorageTemporaryCredentials(ctx).Provider(provider).Path(path).Execute()

获取上传文件临时凭证



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    provider := openapiclient.StorageProvider("qcloud") // StorageProvider | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 (optional)
    path := "path_example" // string | 上传路径，可以为空 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageApi.GetStorageTemporaryCredentials(context.Background()).Provider(provider).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.GetStorageTemporaryCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageTemporaryCredentials`: StorageTemporaryCredentials
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.GetStorageTemporaryCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageTemporaryCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | [**StorageProvider**](StorageProvider.md) | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 | 
 **path** | **string** | 上传路径，可以为空 | 

### Return type

[**StorageTemporaryCredentials**](StorageTemporaryCredentials.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

