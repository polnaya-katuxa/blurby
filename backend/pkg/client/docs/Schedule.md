# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Periodic** | Pointer to **bool** |  | [optional] 
**Finished** | Pointer to **bool** |  | [optional] 
**NextTime** | Pointer to **string** |  | [optional] 
**Span** | **string** |  | 

## Methods

### NewSchedule

`func NewSchedule(span string, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodic

`func (o *Schedule) GetPeriodic() bool`

GetPeriodic returns the Periodic field if non-nil, zero value otherwise.

### GetPeriodicOk

`func (o *Schedule) GetPeriodicOk() (*bool, bool)`

GetPeriodicOk returns a tuple with the Periodic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodic

`func (o *Schedule) SetPeriodic(v bool)`

SetPeriodic sets Periodic field to given value.

### HasPeriodic

`func (o *Schedule) HasPeriodic() bool`

HasPeriodic returns a boolean if a field has been set.

### GetFinished

`func (o *Schedule) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *Schedule) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *Schedule) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *Schedule) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetNextTime

`func (o *Schedule) GetNextTime() string`

GetNextTime returns the NextTime field if non-nil, zero value otherwise.

### GetNextTimeOk

`func (o *Schedule) GetNextTimeOk() (*string, bool)`

GetNextTimeOk returns a tuple with the NextTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTime

`func (o *Schedule) SetNextTime(v string)`

SetNextTime sets NextTime field to given value.

### HasNextTime

`func (o *Schedule) HasNextTime() bool`

HasNextTime returns a boolean if a field has been set.

### GetSpan

`func (o *Schedule) GetSpan() string`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *Schedule) GetSpanOk() (*string, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *Schedule) SetSpan(v string)`

SetSpan sets Span field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


