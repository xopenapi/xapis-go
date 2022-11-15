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

// UpdateConfigGroup struct for UpdateConfigGroup
type UpdateConfigGroup struct {
	// 配置组名称
	Name *string `json:"name,omitempty"`
}

// NewUpdateConfigGroup instantiates a new UpdateConfigGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConfigGroup() *UpdateConfigGroup {
	this := UpdateConfigGroup{}
	return &this
}

// NewUpdateConfigGroupWithDefaults instantiates a new UpdateConfigGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConfigGroupWithDefaults() *UpdateConfigGroup {
	this := UpdateConfigGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateConfigGroup) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConfigGroup) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateConfigGroup) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateConfigGroup) SetName(v string) {
	o.Name = &v
}

func (o UpdateConfigGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateConfigGroup struct {
	value *UpdateConfigGroup
	isSet bool
}

func (v NullableUpdateConfigGroup) Get() *UpdateConfigGroup {
	return v.value
}

func (v *NullableUpdateConfigGroup) Set(val *UpdateConfigGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConfigGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConfigGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConfigGroup(val *UpdateConfigGroup) *NullableUpdateConfigGroup {
	return &NullableUpdateConfigGroup{value: val, isSet: true}
}

func (v NullableUpdateConfigGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConfigGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

