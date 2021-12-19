# StakeAccountDataStake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegation** | Pointer to [**StakeAccountDataStakeDelegation**](StakeAccountDataStakeDelegation.md) |  | [optional] 
**CreditsObserved** | Pointer to **int32** |  | [optional] 

## Methods

### NewStakeAccountDataStake

`func NewStakeAccountDataStake() *StakeAccountDataStake`

NewStakeAccountDataStake instantiates a new StakeAccountDataStake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakeAccountDataStakeWithDefaults

`func NewStakeAccountDataStakeWithDefaults() *StakeAccountDataStake`

NewStakeAccountDataStakeWithDefaults instantiates a new StakeAccountDataStake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegation

`func (o *StakeAccountDataStake) GetDelegation() StakeAccountDataStakeDelegation`

GetDelegation returns the Delegation field if non-nil, zero value otherwise.

### GetDelegationOk

`func (o *StakeAccountDataStake) GetDelegationOk() (*StakeAccountDataStakeDelegation, bool)`

GetDelegationOk returns a tuple with the Delegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegation

`func (o *StakeAccountDataStake) SetDelegation(v StakeAccountDataStakeDelegation)`

SetDelegation sets Delegation field to given value.

### HasDelegation

`func (o *StakeAccountDataStake) HasDelegation() bool`

HasDelegation returns a boolean if a field has been set.

### GetCreditsObserved

`func (o *StakeAccountDataStake) GetCreditsObserved() int32`

GetCreditsObserved returns the CreditsObserved field if non-nil, zero value otherwise.

### GetCreditsObservedOk

`func (o *StakeAccountDataStake) GetCreditsObservedOk() (*int32, bool)`

GetCreditsObservedOk returns a tuple with the CreditsObserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsObserved

`func (o *StakeAccountDataStake) SetCreditsObserved(v int32)`

SetCreditsObserved sets CreditsObserved field to given value.

### HasCreditsObserved

`func (o *StakeAccountDataStake) HasCreditsObserved() bool`

HasCreditsObserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


