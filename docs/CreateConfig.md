# CreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | 所属的资源ID | 
**GroupId** | Pointer to **string** | 所属的配置组ID | [optional] 
**TemplateId** | Pointer to **string** | 配置模板ID | [optional] 
**Key** | **string** | 配置项 | 
**Value** | **map[string]interface{}** | 配置内容 | 

## Methods

### NewCreateConfig

`func NewCreateConfig(resourceId string, key string, value map[string]interface{}, ) *CreateConfig`

NewCreateConfig instantiates a new CreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConfigWithDefaults

`func NewCreateConfigWithDefaults() *CreateConfig`

NewCreateConfigWithDefaults instantiates a new CreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *CreateConfig) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CreateConfig) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CreateConfig) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetGroupId

`func (o *CreateConfig) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateConfig) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateConfig) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CreateConfig) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetTemplateId

`func (o *CreateConfig) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateConfig) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateConfig) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CreateConfig) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetKey

`func (o *CreateConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateConfig) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *CreateConfig) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateConfig) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateConfig) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


