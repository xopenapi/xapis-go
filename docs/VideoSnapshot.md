# VideoSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | [**Provider**](Provider.md) |  | 
**Url** | **string** | 截图的链接地址 | 

## Methods

### NewVideoSnapshot

`func NewVideoSnapshot(provider Provider, url string, ) *VideoSnapshot`

NewVideoSnapshot instantiates a new VideoSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSnapshotWithDefaults

`func NewVideoSnapshotWithDefaults() *VideoSnapshot`

NewVideoSnapshotWithDefaults instantiates a new VideoSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VideoSnapshot) GetProvider() Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VideoSnapshot) GetProviderOk() (*Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VideoSnapshot) SetProvider(v Provider)`

SetProvider sets Provider field to given value.


### GetUrl

`func (o *VideoSnapshot) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoSnapshot) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoSnapshot) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


