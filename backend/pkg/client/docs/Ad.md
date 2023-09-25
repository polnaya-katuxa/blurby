# Ad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**CreateTime** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**UserID** | **string** |  | 
**Schedule** | [**Schedule**](Schedule.md) |  | 

## Methods

### NewAd

`func NewAd(content string, userID string, schedule Schedule, ) *Ad`

NewAd instantiates a new Ad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdWithDefaults

`func NewAdWithDefaults() *Ad`

NewAdWithDefaults instantiates a new Ad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Ad) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Ad) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Ad) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreateTime

`func (o *Ad) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Ad) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Ad) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Ad) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetFilters

`func (o *Ad) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Ad) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Ad) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Ad) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetUserID

`func (o *Ad) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *Ad) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *Ad) SetUserID(v string)`

SetUserID sets UserID field to given value.


### GetSchedule

`func (o *Ad) GetSchedule() Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Ad) GetScheduleOk() (*Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Ad) SetSchedule(v Schedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


