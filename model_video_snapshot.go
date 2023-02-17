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

// VideoSnapshot struct for VideoSnapshot
type VideoSnapshot struct {
	Provider Provider `json:"provider"`
	// 截图的链接地址
	Url string `json:"url"`
}

// NewVideoSnapshot instantiates a new VideoSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoSnapshot(provider Provider, url string) *VideoSnapshot {
	this := VideoSnapshot{}
	this.Provider = provider
	this.Url = url
	return &this
}

// NewVideoSnapshotWithDefaults instantiates a new VideoSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoSnapshotWithDefaults() *VideoSnapshot {
	this := VideoSnapshot{}
	return &this
}

// GetProvider returns the Provider field value
func (o *VideoSnapshot) GetProvider() Provider {
	if o == nil {
		var ret Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *VideoSnapshot) GetProviderOk() (*Provider, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *VideoSnapshot) SetProvider(v Provider) {
	o.Provider = v
}

// GetUrl returns the Url field value
func (o *VideoSnapshot) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VideoSnapshot) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VideoSnapshot) SetUrl(v string) {
	o.Url = v
}

func (o VideoSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableVideoSnapshot struct {
	value *VideoSnapshot
	isSet bool
}

func (v NullableVideoSnapshot) Get() *VideoSnapshot {
	return v.value
}

func (v *NullableVideoSnapshot) Set(val *VideoSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoSnapshot(val *VideoSnapshot) *NullableVideoSnapshot {
	return &NullableVideoSnapshot{value: val, isSet: true}
}

func (v NullableVideoSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


