# SendSMSResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | 是否发送成功 | 
**Provider** | [**Provider**](Provider.md) |  | 
**ProvierRequestId** | Pointer to **string** | 服务商返回的请求ID | [optional] 

## Methods

### NewSendSMSResult

`func NewSendSMSResult(success bool, provider Provider, ) *SendSMSResult`

NewSendSMSResult instantiates a new SendSMSResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendSMSResultWithDefaults

`func NewSendSMSResultWithDefaults() *SendSMSResult`

NewSendSMSResultWithDefaults instantiates a new SendSMSResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SendSMSResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SendSMSResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SendSMSResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetProvider

`func (o *SendSMSResult) GetProvider() Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SendSMSResult) GetProviderOk() (*Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SendSMSResult) SetProvider(v Provider)`

SetProvider sets Provider field to given value.


### GetProvierRequestId

`func (o *SendSMSResult) GetProvierRequestId() string`

GetProvierRequestId returns the ProvierRequestId field if non-nil, zero value otherwise.

### GetProvierRequestIdOk

`func (o *SendSMSResult) GetProvierRequestIdOk() (*string, bool)`

GetProvierRequestIdOk returns a tuple with the ProvierRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvierRequestId

`func (o *SendSMSResult) SetProvierRequestId(v string)`

SetProvierRequestId sets ProvierRequestId field to given value.

### HasProvierRequestId

`func (o *SendSMSResult) HasProvierRequestId() bool`

HasProvierRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


