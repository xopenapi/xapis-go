/*
xapi services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xapis

import (
	"encoding/json"
)

// UpdateConfigTemplate struct for UpdateConfigTemplate
type UpdateConfigTemplate struct {
	// 所属的配置组ID
	GroupId *string `json:"group_id,omitempty"`
	// 配置项
	Key *string `json:"key,omitempty"`
	// 配置内容
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewUpdateConfigTemplate instantiates a new UpdateConfigTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConfigTemplate() *UpdateConfigTemplate {
	this := UpdateConfigTemplate{}
	return &this
}

// NewUpdateConfigTemplateWithDefaults instantiates a new UpdateConfigTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConfigTemplateWithDefaults() *UpdateConfigTemplate {
	this := UpdateConfigTemplate{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *UpdateConfigTemplate) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConfigTemplate) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *UpdateConfigTemplate) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *UpdateConfigTemplate) SetGroupId(v string) {
	o.GroupId = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateConfigTemplate) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConfigTemplate) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateConfigTemplate) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateConfigTemplate) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UpdateConfigTemplate) GetValue() map[string]interface{} {
	if o == nil || isNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConfigTemplate) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Value) {
    return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UpdateConfigTemplate) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *UpdateConfigTemplate) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o UpdateConfigTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateConfigTemplate struct {
	value *UpdateConfigTemplate
	isSet bool
}

func (v NullableUpdateConfigTemplate) Get() *UpdateConfigTemplate {
	return v.value
}

func (v *NullableUpdateConfigTemplate) Set(val *UpdateConfigTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConfigTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConfigTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConfigTemplate(val *UpdateConfigTemplate) *NullableUpdateConfigTemplate {
	return &NullableUpdateConfigTemplate{value: val, isSet: true}
}

func (v NullableUpdateConfigTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConfigTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

