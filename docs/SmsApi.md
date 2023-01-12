# \SmsApi

All URIs are relative to *https://api.growingbox.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendSMS**](SmsApi.md#SendSMS) | **Post** /sms/v1/send | 发送短信



## SendSMS

> SendSMSResult SendSMS(ctx).SendSMS(sendSMS).Execute()

发送短信



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
    sendSMS := *openapiclient.NewSendSMS("18665311790", "1668403") // SendSMS |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmsApi.SendSMS(context.Background()).SendSMS(sendSMS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.SendSMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSMS`: SendSMSResult
    fmt.Fprintf(os.Stdout, "Response from `SmsApi.SendSMS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendSMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendSMS** | [**SendSMS**](SendSMS.md) |  | 

### Return type

[**SendSMSResult**](SendSMSResult.md)

### Authorization

[apiKey](../README.md#apiKey), [bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

