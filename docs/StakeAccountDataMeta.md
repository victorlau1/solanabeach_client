# StakeAccountDataMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RentExemptReserve** | Pointer to **int32** |  | [optional] 
**Authorized** | Pointer to [**StakeAccountDataMetaAuthorized**](StakeAccountDataMetaAuthorized.md) |  | [optional] 

## Methods

### NewStakeAccountDataMeta

`func NewStakeAccountDataMeta() *StakeAccountDataMeta`

NewStakeAccountDataMeta instantiates a new StakeAccountDataMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakeAccountDataMetaWithDefaults

`func NewStakeAccountDataMetaWithDefaults() *StakeAccountDataMeta`

NewStakeAccountDataMetaWithDefaults instantiates a new StakeAccountDataMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRentExemptReserve

`func (o *StakeAccountDataMeta) GetRentExemptReserve() int32`

GetRentExemptReserve returns the RentExemptReserve field if non-nil, zero value otherwise.

### GetRentExemptReserveOk

`func (o *StakeAccountDataMeta) GetRentExemptReserveOk() (*int32, bool)`

GetRentExemptReserveOk returns a tuple with the RentExemptReserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentExemptReserve

`func (o *StakeAccountDataMeta) SetRentExemptReserve(v int32)`

SetRentExemptReserve sets RentExemptReserve field to given value.

### HasRentExemptReserve

`func (o *StakeAccountDataMeta) HasRentExemptReserve() bool`

HasRentExemptReserve returns a boolean if a field has been set.

### GetAuthorized

`func (o *StakeAccountDataMeta) GetAuthorized() StakeAccountDataMetaAuthorized`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *StakeAccountDataMeta) GetAuthorizedOk() (*StakeAccountDataMetaAuthorized, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *StakeAccountDataMeta) SetAuthorized(v StakeAccountDataMetaAuthorized)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *StakeAccountDataMeta) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


