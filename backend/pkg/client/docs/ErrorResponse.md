# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**SystemMessage** | **string** |  | 

## Methods

### NewErrorResponse

`func NewErrorResponse(message string, systemMessage string, ) *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSystemMessage

`func (o *ErrorResponse) GetSystemMessage() string`

GetSystemMessage returns the SystemMessage field if non-nil, zero value otherwise.

### GetSystemMessageOk

`func (o *ErrorResponse) GetSystemMessageOk() (*string, bool)`

GetSystemMessageOk returns a tuple with the SystemMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMessage

`func (o *ErrorResponse) SetSystemMessage(v string)`

SetSystemMessage sets SystemMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


