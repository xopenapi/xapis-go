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

// StorageTemporaryCredentials struct for StorageTemporaryCredentials
type StorageTemporaryCredentials struct {
	Provider StorageProvider `json:"provider"`
	// 存储桶
	Bucket string `json:"bucket"`
	// 地区
	Region string `json:"region"`
	// 存储对应的CDN地址
	Cdn *string `json:"cdn,omitempty"`
	// 上传路径对应的CDN地址
	Url *string `json:"url,omitempty"`
	// 存储平台的临时上传凭证参数
	Credentials map[string]interface{} `json:"credentials"`
}

// NewStorageTemporaryCredentials instantiates a new StorageTemporaryCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageTemporaryCredentials(provider StorageProvider, bucket string, region string, credentials map[string]interface{}) *StorageTemporaryCredentials {
	this := StorageTemporaryCredentials{}
	this.Provider = provider
	this.Bucket = bucket
	this.Region = region
	this.Credentials = credentials
	return &this
}

// NewStorageTemporaryCredentialsWithDefaults instantiates a new StorageTemporaryCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageTemporaryCredentialsWithDefaults() *StorageTemporaryCredentials {
	this := StorageTemporaryCredentials{}
	return &this
}

// GetProvider returns the Provider field value
func (o *StorageTemporaryCredentials) GetProvider() StorageProvider {
	if o == nil {
		var ret StorageProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *StorageTemporaryCredentials) GetProviderOk() (*StorageProvider, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *StorageTemporaryCredentials) SetProvider(v StorageProvider) {
	o.Provider = v
}

// GetBucket returns the Bucket field value
func (o *StorageTemporaryCredentials) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *StorageTemporaryCredentials) GetBucketOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *StorageTemporaryCredentials) SetBucket(v string) {
	o.Bucket = v
}

// GetRegion returns the Region field value
func (o *StorageTemporaryCredentials) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *StorageTemporaryCredentials) GetRegionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *StorageTemporaryCredentials) SetRegion(v string) {
	o.Region = v
}

// GetCdn returns the Cdn field value if set, zero value otherwise.
func (o *StorageTemporaryCredentials) GetCdn() string {
	if o == nil || isNil(o.Cdn) {
		var ret string
		return ret
	}
	return *o.Cdn
}

// GetCdnOk returns a tuple with the Cdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTemporaryCredentials) GetCdnOk() (*string, bool) {
	if o == nil || isNil(o.Cdn) {
    return nil, false
	}
	return o.Cdn, true
}

// HasCdn returns a boolean if a field has been set.
func (o *StorageTemporaryCredentials) HasCdn() bool {
	if o != nil && !isNil(o.Cdn) {
		return true
	}

	return false
}

// SetCdn gets a reference to the given string and assigns it to the Cdn field.
func (o *StorageTemporaryCredentials) SetCdn(v string) {
	o.Cdn = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *StorageTemporaryCredentials) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageTemporaryCredentials) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *StorageTemporaryCredentials) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *StorageTemporaryCredentials) SetUrl(v string) {
	o.Url = &v
}

// GetCredentials returns the Credentials field value
func (o *StorageTemporaryCredentials) GetCredentials() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *StorageTemporaryCredentials) GetCredentialsOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *StorageTemporaryCredentials) SetCredentials(v map[string]interface{}) {
	o.Credentials = v
}

func (o StorageTemporaryCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["bucket"] = o.Bucket
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Cdn) {
		toSerialize["cdn"] = o.Cdn
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	return json.Marshal(toSerialize)
}

type NullableStorageTemporaryCredentials struct {
	value *StorageTemporaryCredentials
	isSet bool
}

func (v NullableStorageTemporaryCredentials) Get() *StorageTemporaryCredentials {
	return v.value
}

func (v *NullableStorageTemporaryCredentials) Set(val *StorageTemporaryCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageTemporaryCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageTemporaryCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageTemporaryCredentials(val *StorageTemporaryCredentials) *NullableStorageTemporaryCredentials {
	return &NullableStorageTemporaryCredentials{value: val, isSet: true}
}

func (v NullableStorageTemporaryCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageTemporaryCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

