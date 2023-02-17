# AudioInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | 音频时长（秒） | 
**Size** | **int32** | 音频大小（字节） | 
**Format** | Pointer to **string** | 音频格式 | [optional] 

## Methods

### NewAudioInfo

`func NewAudioInfo(duration int32, size int32, ) *AudioInfo`

NewAudioInfo instantiates a new AudioInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioInfoWithDefaults

`func NewAudioInfoWithDefaults() *AudioInfo`

NewAudioInfoWithDefaults instantiates a new AudioInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *AudioInfo) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AudioInfo) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AudioInfo) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetSize

`func (o *AudioInfo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AudioInfo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AudioInfo) SetSize(v int32)`

SetSize sets Size field to given value.


### GetFormat

`func (o *AudioInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AudioInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AudioInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AudioInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


