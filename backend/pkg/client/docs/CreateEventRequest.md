# CreateEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**Event**](Event.md) |  | 

## Methods

### NewCreateEventRequest

`func NewCreateEventRequest(event Event, ) *CreateEventRequest`

NewCreateEventRequest instantiates a new CreateEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEventRequestWithDefaults

`func NewCreateEventRequestWithDefaults() *CreateEventRequest`

NewCreateEventRequestWithDefaults instantiates a new CreateEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CreateEventRequest) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CreateEventRequest) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CreateEventRequest) SetEvent(v Event)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


