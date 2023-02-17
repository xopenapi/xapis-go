# VideoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | 视频时长（秒） | 
**Width** | **int32** | 视频宽度（像素） | 
**Height** | **int32** | 视频高度（像素） | 
**Size** | **int32** | 视频大小（字节） | 
**Format** | Pointer to **string** | 视频格式 | [optional] 

## Methods

### NewVideoInfo

`func NewVideoInfo(duration int32, width int32, height int32, size int32, ) *VideoInfo`

NewVideoInfo instantiates a new VideoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoInfoWithDefaults

`func NewVideoInfoWithDefaults() *VideoInfo`

NewVideoInfoWithDefaults instantiates a new VideoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *VideoInfo) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoInfo) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoInfo) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetWidth

`func (o *VideoInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VideoInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VideoInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *VideoInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VideoInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VideoInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetSize

`func (o *VideoInfo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VideoInfo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VideoInfo) SetSize(v int32)`

SetSize sets Size field to given value.


### GetFormat

`func (o *VideoInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *VideoInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *VideoInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *VideoInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


