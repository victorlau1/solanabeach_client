# StakeAccountDataStakeDelegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoterPubkey** | Pointer to **string** |  | [optional] 
**Stake** | Pointer to **int32** |  | [optional] 
**ActivationEpoch** | Pointer to **int32** |  | [optional] 
**DeactivationEpoch** | Pointer to **int32** |  | [optional] 
**WarmupCooldownRate** | Pointer to **int32** |  | [optional] 
**ValidatorInfo** | Pointer to [**BlockProposerData**](BlockProposerData.md) |  | [optional] 

## Methods

### NewStakeAccountDataStakeDelegation

`func NewStakeAccountDataStakeDelegation() *StakeAccountDataStakeDelegation`

NewStakeAccountDataStakeDelegation instantiates a new StakeAccountDataStakeDelegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakeAccountDataStakeDelegationWithDefaults

`func NewStakeAccountDataStakeDelegationWithDefaults() *StakeAccountDataStakeDelegation`

NewStakeAccountDataStakeDelegationWithDefaults instantiates a new StakeAccountDataStakeDelegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoterPubkey

`func (o *StakeAccountDataStakeDelegation) GetVoterPubkey() string`

GetVoterPubkey returns the VoterPubkey field if non-nil, zero value otherwise.

### GetVoterPubkeyOk

`func (o *StakeAccountDataStakeDelegation) GetVoterPubkeyOk() (*string, bool)`

GetVoterPubkeyOk returns a tuple with the VoterPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoterPubkey

`func (o *StakeAccountDataStakeDelegation) SetVoterPubkey(v string)`

SetVoterPubkey sets VoterPubkey field to given value.

### HasVoterPubkey

`func (o *StakeAccountDataStakeDelegation) HasVoterPubkey() bool`

HasVoterPubkey returns a boolean if a field has been set.

### GetStake

`func (o *StakeAccountDataStakeDelegation) GetStake() int32`

GetStake returns the Stake field if non-nil, zero value otherwise.

### GetStakeOk

`func (o *StakeAccountDataStakeDelegation) GetStakeOk() (*int32, bool)`

GetStakeOk returns a tuple with the Stake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStake

`func (o *StakeAccountDataStakeDelegation) SetStake(v int32)`

SetStake sets Stake field to given value.

### HasStake

`func (o *StakeAccountDataStakeDelegation) HasStake() bool`

HasStake returns a boolean if a field has been set.

### GetActivationEpoch

`func (o *StakeAccountDataStakeDelegation) GetActivationEpoch() int32`

GetActivationEpoch returns the ActivationEpoch field if non-nil, zero value otherwise.

### GetActivationEpochOk

`func (o *StakeAccountDataStakeDelegation) GetActivationEpochOk() (*int32, bool)`

GetActivationEpochOk returns a tuple with the ActivationEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationEpoch

`func (o *StakeAccountDataStakeDelegation) SetActivationEpoch(v int32)`

SetActivationEpoch sets ActivationEpoch field to given value.

### HasActivationEpoch

`func (o *StakeAccountDataStakeDelegation) HasActivationEpoch() bool`

HasActivationEpoch returns a boolean if a field has been set.

### GetDeactivationEpoch

`func (o *StakeAccountDataStakeDelegation) GetDeactivationEpoch() int32`

GetDeactivationEpoch returns the DeactivationEpoch field if non-nil, zero value otherwise.

### GetDeactivationEpochOk

`func (o *StakeAccountDataStakeDelegation) GetDeactivationEpochOk() (*int32, bool)`

GetDeactivationEpochOk returns a tuple with the DeactivationEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivationEpoch

`func (o *StakeAccountDataStakeDelegation) SetDeactivationEpoch(v int32)`

SetDeactivationEpoch sets DeactivationEpoch field to given value.

### HasDeactivationEpoch

`func (o *StakeAccountDataStakeDelegation) HasDeactivationEpoch() bool`

HasDeactivationEpoch returns a boolean if a field has been set.

### GetWarmupCooldownRate

`func (o *StakeAccountDataStakeDelegation) GetWarmupCooldownRate() int32`

GetWarmupCooldownRate returns the WarmupCooldownRate field if non-nil, zero value otherwise.

### GetWarmupCooldownRateOk

`func (o *StakeAccountDataStakeDelegation) GetWarmupCooldownRateOk() (*int32, bool)`

GetWarmupCooldownRateOk returns a tuple with the WarmupCooldownRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupCooldownRate

`func (o *StakeAccountDataStakeDelegation) SetWarmupCooldownRate(v int32)`

SetWarmupCooldownRate sets WarmupCooldownRate field to given value.

### HasWarmupCooldownRate

`func (o *StakeAccountDataStakeDelegation) HasWarmupCooldownRate() bool`

HasWarmupCooldownRate returns a boolean if a field has been set.

### GetValidatorInfo

`func (o *StakeAccountDataStakeDelegation) GetValidatorInfo() BlockProposerData`

GetValidatorInfo returns the ValidatorInfo field if non-nil, zero value otherwise.

### GetValidatorInfoOk

`func (o *StakeAccountDataStakeDelegation) GetValidatorInfoOk() (*BlockProposerData, bool)`

GetValidatorInfoOk returns a tuple with the ValidatorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorInfo

`func (o *StakeAccountDataStakeDelegation) SetValidatorInfo(v BlockProposerData)`

SetValidatorInfo sets ValidatorInfo field to given value.

### HasValidatorInfo

`func (o *StakeAccountDataStakeDelegation) HasValidatorInfo() bool`

HasValidatorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


