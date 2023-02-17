# \SilkCodecApi

All URIs are relative to *https://api.growingbox.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DecodeSilkV3**](SilkCodecApi.md#DecodeSilkV3) | **Post** /silk/v3/decode | 解码silk-v3格式的语音
[**DecodeSilkV3ByURL**](SilkCodecApi.md#DecodeSilkV3ByURL) | **Get** /silk/v3/decode | 将链接解码silk-v3格式的语音
[**EncodeSilkV3**](SilkCodecApi.md#EncodeSilkV3) | **Post** /silk/v3/encode | 将语音编码为silk-v3格式
[**EncodeSilkV3ByURL**](SilkCodecApi.md#EncodeSilkV3ByURL) | **Get** /silk/v3/encode | 将语音链接编码为silk-v3格式



## DecodeSilkV3

> *os.File DecodeSilkV3(ctx).Format(format).File(file).Execute()

解码silk-v3格式的语音



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
    format := "format_example" // string | 解码后的音频格式，如：mp3、m4a、mov等。 (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SilkCodecApi.DecodeSilkV3(context.Background()).Format(format).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SilkCodecApi.DecodeSilkV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecodeSilkV3`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SilkCodecApi.DecodeSilkV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecodeSilkV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | 解码后的音频格式，如：mp3、m4a、mov等。 | 
 **file** | ***os.File** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecodeSilkV3ByURL

> *os.File DecodeSilkV3ByURL(ctx).Url(url).Format(format).Execute()

将链接解码silk-v3格式的语音



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
    url := "url_example" // string | 语音链接
    format := "format_example" // string | 解码后的音频格式，如：mp3、m4a、mov等。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SilkCodecApi.DecodeSilkV3ByURL(context.Background()).Url(url).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SilkCodecApi.DecodeSilkV3ByURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecodeSilkV3ByURL`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SilkCodecApi.DecodeSilkV3ByURL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecodeSilkV3ByURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | 语音链接 | 
 **format** | **string** | 解码后的音频格式，如：mp3、m4a、mov等。 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncodeSilkV3

> *os.File EncodeSilkV3(ctx).File(file).Execute()

将语音编码为silk-v3格式



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
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SilkCodecApi.EncodeSilkV3(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SilkCodecApi.EncodeSilkV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncodeSilkV3`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SilkCodecApi.EncodeSilkV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEncodeSilkV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncodeSilkV3ByURL

> *os.File EncodeSilkV3ByURL(ctx).Url(url).Execute()

将语音链接编码为silk-v3格式



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
    url := "url_example" // string | 语音链接

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SilkCodecApi.EncodeSilkV3ByURL(context.Background()).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SilkCodecApi.EncodeSilkV3ByURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncodeSilkV3ByURL`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SilkCodecApi.EncodeSilkV3ByURL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEncodeSilkV3ByURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | 语音链接 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

