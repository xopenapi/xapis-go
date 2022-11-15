# Go API client for xapis

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 0.0.3
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import xapis "github.com/xopenapi/xapis-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), xapis.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), xapis.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), xapis.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), xapis.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.growingbox.cn*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConfigApi* | [**CreateConfig**](docs/ConfigApi.md#createconfig) | **Post** /config/v1/configs | 创建配置
*ConfigApi* | [**CreateConfigGroup**](docs/ConfigApi.md#createconfiggroup) | **Post** /config/v1/config_groups | 创建配置组
*ConfigApi* | [**CreateConfigTemplate**](docs/ConfigApi.md#createconfigtemplate) | **Post** /config/v1/config_templates | 创建配置模版
*ConfigApi* | [**DeleteConfig**](docs/ConfigApi.md#deleteconfig) | **Delete** /config/v1/configs/{id} | 删除配置
*ConfigApi* | [**DeleteConfigGroup**](docs/ConfigApi.md#deleteconfiggroup) | **Delete** /config/v1/config_groups/{id} | 删除配置组
*ConfigApi* | [**DeleteConfigTemplate**](docs/ConfigApi.md#deleteconfigtemplate) | **Delete** /config/v1/config_templates/{id} | 删除配置模板
*ConfigApi* | [**GetConfig**](docs/ConfigApi.md#getconfig) | **Get** /config/v1/configs/{id} | 查询配置
*ConfigApi* | [**GetConfigByKey**](docs/ConfigApi.md#getconfigbykey) | **Get** /config/v1/configs | 根据配置项查询配置
*ConfigApi* | [**GetConfigGroup**](docs/ConfigApi.md#getconfiggroup) | **Get** /config/v1/config_groups/{id} | 查询配置组
*ConfigApi* | [**GetConfigTemplate**](docs/ConfigApi.md#getconfigtemplate) | **Get** /config/v1/config_templates/{id} | 查询配置模板
*ConfigApi* | [**GetConfigTemplateByKey**](docs/ConfigApi.md#getconfigtemplatebykey) | **Get** /config/v1/config_templates | 根据配置项查询配置模版
*ConfigApi* | [**UpdateConfig**](docs/ConfigApi.md#updateconfig) | **Patch** /config/v1/configs/{id} | 更新配置
*ConfigApi* | [**UpdateConfigGroup**](docs/ConfigApi.md#updateconfiggroup) | **Patch** /config/v1/config_groups/{id} | 更新配置组
*ConfigApi* | [**UpdateConfigTemplate**](docs/ConfigApi.md#updateconfigtemplate) | **Patch** /config/v1/config_templates/{id} | 更新配置模板
*StorageApi* | [**GetStorageTemporaryCredentials**](docs/StorageApi.md#getstoragetemporarycredentials) | **Get** /storage/v1/temporary_credentials | 获取上传文件临时凭证


## Documentation For Models

 - [ApiError](docs/ApiError.md)
 - [Config](docs/Config.md)
 - [ConfigGroup](docs/ConfigGroup.md)
 - [ConfigTemplate](docs/ConfigTemplate.md)
 - [CreateConfig](docs/CreateConfig.md)
 - [CreateConfigGroup](docs/CreateConfigGroup.md)
 - [CreateConfigTemplate](docs/CreateConfigTemplate.md)
 - [StorageProvider](docs/StorageProvider.md)
 - [StorageTemporaryCredentials](docs/StorageTemporaryCredentials.md)
 - [UpdateConfig](docs/UpdateConfig.md)
 - [UpdateConfigGroup](docs/UpdateConfigGroup.md)
 - [UpdateConfigTemplate](docs/UpdateConfigTemplate.md)


## Documentation For Authorization



### bearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


### oauth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **configs:read**: 读取配置、配置模版、配置组
 - **configs:write**: 创建、更新配置、配置模版、配置组
 - **configs:delete**: 删除配置、配置模版、配置组
 - **storage.credentials:read**: 获取云存储上传凭证

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



