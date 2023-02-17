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

// SendSMSResult struct for SendSMSResult
type SendSMSResult struct {
	// 是否发送成功
	Success bool `json:"success"`
	Provider Provider `json:"provider"`
	// 服务商返回的请求ID
	ProvierRequestId *string `json:"provier_request_id,omitempty"`
}

// NewSendSMSResult instantiates a new SendSMSResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendSMSResult(success bool, provider Provider) *SendSMSResult {
	this := SendSMSResult{}
	this.Success = success
	this.Provider = provider
	return &this
}

// NewSendSMSResultWithDefaults instantiates a new SendSMSResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendSMSResultWithDefaults() *SendSMSResult {
	this := SendSMSResult{}
	return &this
}

// GetSuccess returns the Success field value
func (o *SendSMSResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SendSMSResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *SendSMSResult) SetSuccess(v bool) {
	o.Success = v
}

// GetProvider returns the Provider field value
func (o *SendSMSResult) GetProvider() Provider {
	if o == nil {
		var ret Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SendSMSResult) GetProviderOk() (*Provider, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SendSMSResult) SetProvider(v Provider) {
	o.Provider = v
}

// GetProvierRequestId returns the ProvierRequestId field value if set, zero value otherwise.
func (o *SendSMSResult) GetProvierRequestId() string {
	if o == nil || isNil(o.ProvierRequestId) {
		var ret string
		return ret
	}
	return *o.ProvierRequestId
}

// GetProvierRequestIdOk returns a tuple with the ProvierRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendSMSResult) GetProvierRequestIdOk() (*string, bool) {
	if o == nil || isNil(o.ProvierRequestId) {
    return nil, false
	}
	return o.ProvierRequestId, true
}

// HasProvierRequestId returns a boolean if a field has been set.
func (o *SendSMSResult) HasProvierRequestId() bool {
	if o != nil && !isNil(o.ProvierRequestId) {
		return true
	}

	return false
}

// SetProvierRequestId gets a reference to the given string and assigns it to the ProvierRequestId field.
func (o *SendSMSResult) SetProvierRequestId(v string) {
	o.ProvierRequestId = &v
}

func (o SendSMSResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["success"] = o.Success
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.ProvierRequestId) {
		toSerialize["provier_request_id"] = o.ProvierRequestId
	}
	return json.Marshal(toSerialize)
}

type NullableSendSMSResult struct {
	value *SendSMSResult
	isSet bool
}

func (v NullableSendSMSResult) Get() *SendSMSResult {
	return v.value
}

func (v *NullableSendSMSResult) Set(val *SendSMSResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSendSMSResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSendSMSResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendSMSResult(val *SendSMSResult) *NullableSendSMSResult {
	return &NullableSendSMSResult{value: val, isSet: true}
}

func (v NullableSendSMSResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendSMSResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


