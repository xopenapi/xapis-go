# UpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **string** | 所属的资源ID | [optional] 
**GroupId** | Pointer to **string** | 所属的配置组ID | [optional] 
**TemplateId** | Pointer to **string** | 配置模板ID | [optional] 
**Key** | Pointer to **string** | 配置项 | [optional] 
**Value** | Pointer to **map[string]interface{}** | 配置内容 | [optional] 

## Methods

### NewUpdateConfig

`func NewUpdateConfig() *UpdateConfig`

NewUpdateConfig instantiates a new UpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConfigWithDefaults

`func NewUpdateConfigWithDefaults() *UpdateConfig`

NewUpdateConfigWithDefaults instantiates a new UpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *UpdateConfig) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *UpdateConfig) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *UpdateConfig) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *UpdateConfig) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetGroupId

`func (o *UpdateConfig) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateConfig) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateConfig) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UpdateConfig) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetTemplateId

`func (o *UpdateConfig) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UpdateConfig) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UpdateConfig) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *UpdateConfig) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetKey

`func (o *UpdateConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *UpdateConfig) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateConfig) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateConfig) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


