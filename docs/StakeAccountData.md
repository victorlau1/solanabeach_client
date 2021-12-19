# StakeAccountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to [**StakeAccountDataMeta**](StakeAccountDataMeta.md) |  | [optional] 
**Lockup** | Pointer to [**StakeAccountDataLockup**](StakeAccountDataLockup.md) |  | [optional] 
**Stake** | Pointer to [**StakeAccountDataStake**](StakeAccountDataStake.md) |  | [optional] 

## Methods

### NewStakeAccountData

`func NewStakeAccountData() *StakeAccountData`

NewStakeAccountData instantiates a new StakeAccountData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakeAccountDataWithDefaults

`func NewStakeAccountDataWithDefaults() *StakeAccountData`

NewStakeAccountDataWithDefaults instantiates a new StakeAccountData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *StakeAccountData) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StakeAccountData) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StakeAccountData) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *StakeAccountData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMeta

`func (o *StakeAccountData) GetMeta() StakeAccountDataMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StakeAccountData) GetMetaOk() (*StakeAccountDataMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StakeAccountData) SetMeta(v StakeAccountDataMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StakeAccountData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLockup

`func (o *StakeAccountData) GetLockup() StakeAccountDataLockup`

GetLockup returns the Lockup field if non-nil, zero value otherwise.

### GetLockupOk

`func (o *StakeAccountData) GetLockupOk() (*StakeAccountDataLockup, bool)`

GetLockupOk returns a tuple with the Lockup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockup

`func (o *StakeAccountData) SetLockup(v StakeAccountDataLockup)`

SetLockup sets Lockup field to given value.

### HasLockup

`func (o *StakeAccountData) HasLockup() bool`

HasLockup returns a boolean if a field has been set.

### GetStake

`func (o *StakeAccountData) GetStake() StakeAccountDataStake`

GetStake returns the Stake field if non-nil, zero value otherwise.

### GetStakeOk

`func (o *StakeAccountData) GetStakeOk() (*StakeAccountDataStake, bool)`

GetStakeOk returns a tuple with the Stake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStake

`func (o *StakeAccountData) SetStake(v StakeAccountDataStake)`

SetStake sets Stake field to given value.

### HasStake

`func (o *StakeAccountData) HasStake() bool`

HasStake returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


