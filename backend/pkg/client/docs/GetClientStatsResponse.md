# GetClientStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientStats** | [**ClientStats**](ClientStats.md) |  | 
**AdStats** | [**[]AdStat**](AdStat.md) |  | 

## Methods

### NewGetClientStatsResponse

`func NewGetClientStatsResponse(clientStats ClientStats, adStats []AdStat, ) *GetClientStatsResponse`

NewGetClientStatsResponse instantiates a new GetClientStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClientStatsResponseWithDefaults

`func NewGetClientStatsResponseWithDefaults() *GetClientStatsResponse`

NewGetClientStatsResponseWithDefaults instantiates a new GetClientStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientStats

`func (o *GetClientStatsResponse) GetClientStats() ClientStats`

GetClientStats returns the ClientStats field if non-nil, zero value otherwise.

### GetClientStatsOk

`func (o *GetClientStatsResponse) GetClientStatsOk() (*ClientStats, bool)`

GetClientStatsOk returns a tuple with the ClientStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientStats

`func (o *GetClientStatsResponse) SetClientStats(v ClientStats)`

SetClientStats sets ClientStats field to given value.


### GetAdStats

`func (o *GetClientStatsResponse) GetAdStats() []AdStat`

GetAdStats returns the AdStats field if non-nil, zero value otherwise.

### GetAdStatsOk

`func (o *GetClientStatsResponse) GetAdStatsOk() (*[]AdStat, bool)`

GetAdStatsOk returns a tuple with the AdStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdStats

`func (o *GetClientStatsResponse) SetAdStats(v []AdStat)`

SetAdStats sets AdStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


