# \ConfigApi

All URIs are relative to *https://api.growingbox.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfig**](ConfigApi.md#CreateConfig) | **Post** /config/v1/configs | 创建配置
[**CreateConfigGroup**](ConfigApi.md#CreateConfigGroup) | **Post** /config/v1/config_groups | 创建配置组
[**CreateConfigTemplate**](ConfigApi.md#CreateConfigTemplate) | **Post** /config/v1/config_templates | 创建配置模版
[**DeleteConfig**](ConfigApi.md#DeleteConfig) | **Delete** /config/v1/configs/{id} | 删除配置
[**DeleteConfigGroup**](ConfigApi.md#DeleteConfigGroup) | **Delete** /config/v1/config_groups/{id} | 删除配置组
[**DeleteConfigTemplate**](ConfigApi.md#DeleteConfigTemplate) | **Delete** /config/v1/config_templates/{id} | 删除配置模板
[**GetConfig**](ConfigApi.md#GetConfig) | **Get** /config/v1/configs/{id} | 查询配置
[**GetConfigByKey**](ConfigApi.md#GetConfigByKey) | **Get** /config/v1/configs | 根据配置项查询配置
[**GetConfigGroup**](ConfigApi.md#GetConfigGroup) | **Get** /config/v1/config_groups/{id} | 查询配置组
[**GetConfigTemplate**](ConfigApi.md#GetConfigTemplate) | **Get** /config/v1/config_templates/{id} | 查询配置模板
[**GetConfigTemplateByKey**](ConfigApi.md#GetConfigTemplateByKey) | **Get** /config/v1/config_templates | 根据配置项查询配置模版
[**UpdateConfig**](ConfigApi.md#UpdateConfig) | **Patch** /config/v1/configs/{id} | 更新配置
[**UpdateConfigGroup**](ConfigApi.md#UpdateConfigGroup) | **Patch** /config/v1/config_groups/{id} | 更新配置组
[**UpdateConfigTemplate**](ConfigApi.md#UpdateConfigTemplate) | **Patch** /config/v1/config_templates/{id} | 更新配置模板



## CreateConfig

> Config CreateConfig(ctx).CreateConfig(createConfig).Execute()

创建配置

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
    createConfig := *openapiclient.NewCreateConfig("ResourceId_example", "Key_example", map[string]interface{}(123)) // CreateConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.CreateConfig(context.Background()).CreateConfig(createConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfig`: Config
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.CreateConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfig** | [**CreateConfig**](CreateConfig.md) |  | 

### Return type

[**Config**](Config.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfigGroup

> ConfigGroup CreateConfigGroup(ctx).CreateConfigGroup(createConfigGroup).Execute()

创建配置组

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
    createConfigGroup := *openapiclient.NewCreateConfigGroup("Name_example") // CreateConfigGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.CreateConfigGroup(context.Background()).CreateConfigGroup(createConfigGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateConfigGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfigGroup`: ConfigGroup
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.CreateConfigGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfigGroup** | [**CreateConfigGroup**](CreateConfigGroup.md) |  | 

### Return type

[**ConfigGroup**](ConfigGroup.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfigTemplate

> ConfigTemplate CreateConfigTemplate(ctx).CreateConfigTemplate(createConfigTemplate).Execute()

创建配置模版

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
    createConfigTemplate := *openapiclient.NewCreateConfigTemplate("Key_example", map[string]interface{}(123)) // CreateConfigTemplate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.CreateConfigTemplate(context.Background()).CreateConfigTemplate(createConfigTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.CreateConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfigTemplate`: ConfigTemplate
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.CreateConfigTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfigTemplate** | [**CreateConfigTemplate**](CreateConfigTemplate.md) |  | 

### Return type

[**ConfigTemplate**](ConfigTemplate.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfig

> DeleteConfig(ctx, id).Execute()

删除配置

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
    id := "id_example" // string | 配置ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.DeleteConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.DeleteConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 配置ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfigGroup

> DeleteConfigGroup(ctx, id).Execute()

删除配置组

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
    id := "id_example" // string | 配置组ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.DeleteConfigGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.DeleteConfigGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 配置组ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfigTemplate

> DeleteConfigTemplate(ctx, id).Execute()

删除配置模板

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
    id := "id_example" // string | 配置模板ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.DeleteConfigTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.DeleteConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 配置模板ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfig

> Config GetConfig(ctx, id).Execute()

查询配置

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
    id := "id_example" // string | 配置ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfig`: Config
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 配置ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Config**](Config.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigByKey

> Config GetConfigByKey(ctx).ResourceId(resourceId).Key(key).Execute()

根据配置项查询配置

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
    resourceId := "resourceId_example" // string | 配置所属的资源ID
    key := "key_example" // string | 配置项

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetConfigByKey(context.Background()).ResourceId(resourceId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfigByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigByKey`: Config
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfigByKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceId** | **string** | 配置所属的资源ID | 
 **key** | **string** | 配置项 | 

### Return type

[**Config**](Config.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigGroup

> ConfigGroup GetConfigGroup(ctx, id).Execute()

查询配置组

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
    id := "id_example" // string | 配置组ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetConfigGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfigGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigGroup`: ConfigGroup
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfigGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 配置组ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigGroup**](ConfigGroup.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigTemplate

> ConfigTemplate GetConfigTemplate(ctx, id).Execute()

查询配置模板

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
    id := "id_example" // string | 配置模板ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetConfigTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigTemplate`: ConfigTemplate
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 配置模板ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigTemplate**](ConfigTemplate.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigTemplateByKey

> ConfigTemplate GetConfigTemplateByKey(ctx).Key(key).Execute()

根据配置项查询配置模版

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
    key := "key_example" // string | 配置项

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetConfigTemplateByKey(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetConfigTemplateByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigTemplateByKey`: ConfigTemplate
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetConfigTemplateByKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigTemplateByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | 配置项 | 

### Return type

[**ConfigTemplate**](ConfigTemplate.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfig

> Config UpdateConfig(ctx, id).UpdateConfig(updateConfig).Execute()

更新配置

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
    id := "id_example" // string | 配置ID
    updateConfig := *openapiclient.NewUpdateConfig() // UpdateConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.UpdateConfig(context.Background(), id).UpdateConfig(updateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfig`: Config
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.UpdateConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 配置ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConfig** | [**UpdateConfig**](UpdateConfig.md) |  | 

### Return type

[**Config**](Config.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigGroup

> ConfigGroup UpdateConfigGroup(ctx, id).UpdateConfigGroup(updateConfigGroup).Execute()

更新配置组

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
    id := "id_example" // string | 配置组ID
    updateConfigGroup := *openapiclient.NewUpdateConfigGroup() // UpdateConfigGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.UpdateConfigGroup(context.Background(), id).UpdateConfigGroup(updateConfigGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateConfigGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigGroup`: ConfigGroup
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.UpdateConfigGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 配置组ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConfigGroup** | [**UpdateConfigGroup**](UpdateConfigGroup.md) |  | 

### Return type

[**ConfigGroup**](ConfigGroup.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigTemplate

> ConfigTemplate UpdateConfigTemplate(ctx, id).UpdateConfigTemplate(updateConfigTemplate).Execute()

更新配置模板

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
    id := "id_example" // string | 配置模板ID
    updateConfigTemplate := *openapiclient.NewUpdateConfigTemplate() // UpdateConfigTemplate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.UpdateConfigTemplate(context.Background(), id).UpdateConfigTemplate(updateConfigTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.UpdateConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigTemplate`: ConfigTemplate
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.UpdateConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 配置模板ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConfigTemplate** | [**UpdateConfigTemplate**](UpdateConfigTemplate.md) |  | 

### Return type

[**ConfigTemplate**](ConfigTemplate.md)

### Authorization

[bearer](../README.md#bearer), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

