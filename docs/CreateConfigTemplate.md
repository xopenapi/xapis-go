# CreateConfigTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | 所属的配置组ID | [optional] 
**Key** | **string** | 配置项 | 
**Value** | **map[string]interface{}** | 配置内容 | 

## Methods

### NewCreateConfigTemplate

`func NewCreateConfigTemplate(key string, value map[string]interface{}, ) *CreateConfigTemplate`

NewCreateConfigTemplate instantiates a new CreateConfigTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConfigTemplateWithDefaults

`func NewCreateConfigTemplateWithDefaults() *CreateConfigTemplate`

NewCreateConfigTemplateWithDefaults instantiates a new CreateConfigTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *CreateConfigTemplate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateConfigTemplate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateConfigTemplate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CreateConfigTemplate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetKey

`func (o *CreateConfigTemplate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateConfigTemplate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateConfigTemplate) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *CreateConfigTemplate) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateConfigTemplate) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateConfigTemplate) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


