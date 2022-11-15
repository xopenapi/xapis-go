# StorageTemporaryCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | [**StorageProvider**](StorageProvider.md) |  | 
**Bucket** | **string** | 存储桶 | 
**Region** | **string** | 地区 | 
**Cdn** | Pointer to **string** | 存储对应的CDN地址 | [optional] 
**Url** | Pointer to **string** | 上传路径对应的CDN地址 | [optional] 
**Credentials** | **map[string]interface{}** | 存储平台的临时上传凭证参数 | 

## Methods

### NewStorageTemporaryCredentials

`func NewStorageTemporaryCredentials(provider StorageProvider, bucket string, region string, credentials map[string]interface{}, ) *StorageTemporaryCredentials`

NewStorageTemporaryCredentials instantiates a new StorageTemporaryCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageTemporaryCredentialsWithDefaults

`func NewStorageTemporaryCredentialsWithDefaults() *StorageTemporaryCredentials`

NewStorageTemporaryCredentialsWithDefaults instantiates a new StorageTemporaryCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *StorageTemporaryCredentials) GetProvider() StorageProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *StorageTemporaryCredentials) GetProviderOk() (*StorageProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *StorageTemporaryCredentials) SetProvider(v StorageProvider)`

SetProvider sets Provider field to given value.


### GetBucket

`func (o *StorageTemporaryCredentials) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *StorageTemporaryCredentials) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *StorageTemporaryCredentials) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetRegion

`func (o *StorageTemporaryCredentials) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageTemporaryCredentials) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageTemporaryCredentials) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCdn

`func (o *StorageTemporaryCredentials) GetCdn() string`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *StorageTemporaryCredentials) GetCdnOk() (*string, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *StorageTemporaryCredentials) SetCdn(v string)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *StorageTemporaryCredentials) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetUrl

`func (o *StorageTemporaryCredentials) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StorageTemporaryCredentials) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StorageTemporaryCredentials) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StorageTemporaryCredentials) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCredentials

`func (o *StorageTemporaryCredentials) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *StorageTemporaryCredentials) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *StorageTemporaryCredentials) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


