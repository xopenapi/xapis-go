# \StorageApi

All URIs are relative to *https://api.growingbox.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudioInfo**](StorageApi.md#GetAudioInfo) | **Get** /storage/v1/audio_info | 获取音频信息
[**GetStorageTemporaryCredentials**](StorageApi.md#GetStorageTemporaryCredentials) | **Get** /storage/v1/temporary_credentials | 获取上传文件临时凭证
[**GetVideoInfo**](StorageApi.md#GetVideoInfo) | **Get** /storage/v1/video_info | 获取视频信息
[**GetVideoSnapshot**](StorageApi.md#GetVideoSnapshot) | **Get** /storage/v1/video_snapshot | 获取视频截图



## GetAudioInfo

> AudioInfo GetAudioInfo(ctx).Path(path).Provider(provider).Execute()

获取音频信息



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
    path := "path_example" // string | 音频在云存储的路径
    provider := openapiclient.Provider("qcloud") // Provider | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageApi.GetAudioInfo(context.Background()).Path(path).Provider(provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.GetAudioInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudioInfo`: AudioInfo
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.GetAudioInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | 音频在云存储的路径 | 
 **provider** | [**Provider**](Provider.md) | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 | 

### Return type

[**AudioInfo**](AudioInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    provider := openapiclient.Provider("qcloud") // Provider | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 (optional)
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
 **provider** | [**Provider**](Provider.md) | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 | 
 **path** | **string** | 上传路径，可以为空 | 

### Return type

[**StorageTemporaryCredentials**](StorageTemporaryCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoInfo

> VideoInfo GetVideoInfo(ctx).Path(path).Provider(provider).Execute()

获取视频信息



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
    path := "path_example" // string | 视频在云存储的路径
    provider := openapiclient.Provider("qcloud") // Provider | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageApi.GetVideoInfo(context.Background()).Path(path).Provider(provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.GetVideoInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVideoInfo`: VideoInfo
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.GetVideoInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | 视频在云存储的路径 | 
 **provider** | [**Provider**](Provider.md) | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 | 

### Return type

[**VideoInfo**](VideoInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoSnapshot

> VideoSnapshot GetVideoSnapshot(ctx).Path(path).Provider(provider).Time(time).SnapshotPath(snapshotPath).Execute()

获取视频截图



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
    path := "path_example" // string | 视频在云存储的路径
    provider := openapiclient.Provider("qcloud") // Provider | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 (optional)
    time := float32(3.4) // float32 | 截图的时间，单位：秒 (optional)
    snapshotPath := "snapshotPath_example" // string | 截图保存在云存储的路径 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageApi.GetVideoSnapshot(context.Background()).Path(path).Provider(provider).Time(time).SnapshotPath(snapshotPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.GetVideoSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVideoSnapshot`: VideoSnapshot
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.GetVideoSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | 视频在云存储的路径 | 
 **provider** | [**Provider**](Provider.md) | 指定使用的云储存平台，可选值有：qcloud（腾讯云）、aliyun（阿里云）。如果未指定，使用默认平台。 | 
 **time** | **float32** | 截图的时间，单位：秒 | 
 **snapshotPath** | **string** | 截图保存在云存储的路径 | 

### Return type

[**VideoSnapshot**](VideoSnapshot.md)

### Authorization

[apiKey](../README.md#apiKey), [bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

