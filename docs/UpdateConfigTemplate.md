# UpdateConfigTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | 所属的配置组ID | [optional] 
**Key** | Pointer to **string** | 配置项 | [optional] 
**Value** | Pointer to **map[string]interface{}** | 配置内容 | [optional] 

## Methods

### NewUpdateConfigTemplate

`func NewUpdateConfigTemplate() *UpdateConfigTemplate`

NewUpdateConfigTemplate instantiates a new UpdateConfigTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConfigTemplateWithDefaults

`func NewUpdateConfigTemplateWithDefaults() *UpdateConfigTemplate`

NewUpdateConfigTemplateWithDefaults instantiates a new UpdateConfigTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UpdateConfigTemplate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateConfigTemplate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateConfigTemplate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UpdateConfigTemplate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetKey

`func (o *UpdateConfigTemplate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateConfigTemplate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateConfigTemplate) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateConfigTemplate) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *UpdateConfigTemplate) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateConfigTemplate) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateConfigTemplate) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateConfigTemplate) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


