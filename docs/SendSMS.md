# SendSMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** | 手机号 | 
**Content** | **string** | 短信内容或模版ID | 
**Params** | Pointer to **map[string]interface{}** | 模版参数 | [optional] 
**Sign** | Pointer to **string** | 短信签名 | [optional] 
**Provider** | Pointer to [**SMSProvider**](SMSProvider.md) |  | [optional] 

## Methods

### NewSendSMS

`func NewSendSMS(phone string, content string, ) *SendSMS`

NewSendSMS instantiates a new SendSMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendSMSWithDefaults

`func NewSendSMSWithDefaults() *SendSMS`

NewSendSMSWithDefaults instantiates a new SendSMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendSMS) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendSMS) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendSMS) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetContent

`func (o *SendSMS) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SendSMS) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SendSMS) SetContent(v string)`

SetContent sets Content field to given value.


### GetParams

`func (o *SendSMS) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SendSMS) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SendSMS) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *SendSMS) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSign

`func (o *SendSMS) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *SendSMS) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *SendSMS) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *SendSMS) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetProvider

`func (o *SendSMS) GetProvider() SMSProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SendSMS) GetProviderOk() (*SMSProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SendSMS) SetProvider(v SMSProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SendSMS) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


