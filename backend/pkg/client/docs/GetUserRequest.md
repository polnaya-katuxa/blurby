# GetUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** |  | 

## Methods

### NewGetUserRequest

`func NewGetUserRequest(login string, ) *GetUserRequest`

NewGetUserRequest instantiates a new GetUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserRequestWithDefaults

`func NewGetUserRequestWithDefaults() *GetUserRequest`

NewGetUserRequestWithDefaults instantiates a new GetUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *GetUserRequest) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *GetUserRequest) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *GetUserRequest) SetLogin(v string)`

SetLogin sets Login field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


