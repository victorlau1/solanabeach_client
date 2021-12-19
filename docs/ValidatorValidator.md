# ValidatorValidator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatedStake** | Pointer to **int32** |  | [optional] 
**StakePercentage** | Pointer to **int32** |  | [optional] 
**Commission** | Pointer to **int32** |  | [optional] 
**EpochCredits** | Pointer to **[]int32** |  | [optional] 
**EpochVoteAccount** | Pointer to **bool** |  | [optional] 
**LastVote** | Pointer to **int32** |  | [optional] 
**NodePubkey** | Pointer to **string** |  | [optional] 
**RootSlot** | Pointer to **int32** |  | [optional] 
**VotePubkey** | Pointer to **string** |  | [optional] 
**BlockProduction** | Pointer to [**ValidatorValidatorBlockProduction**](ValidatorValidatorBlockProduction.md) |  | [optional] 
**DelegatingStakeAccounts** | Pointer to [**[]StakeAccount**](StakeAccount.md) |  | [optional] 
**DelegatorCount** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to [**NonvalidatorsLocation**](NonvalidatorsLocation.md) |  | [optional] 
**Moniker** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**PictureURL** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Asn** | Pointer to [**ValidatorValidatorAsn**](ValidatorValidatorAsn.md) |  | [optional] 

## Methods

### NewValidatorValidator

`func NewValidatorValidator() *ValidatorValidator`

NewValidatorValidator instantiates a new ValidatorValidator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorValidatorWithDefaults

`func NewValidatorValidatorWithDefaults() *ValidatorValidator`

NewValidatorValidatorWithDefaults instantiates a new ValidatorValidator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivatedStake

`func (o *ValidatorValidator) GetActivatedStake() int32`

GetActivatedStake returns the ActivatedStake field if non-nil, zero value otherwise.

### GetActivatedStakeOk

`func (o *ValidatorValidator) GetActivatedStakeOk() (*int32, bool)`

GetActivatedStakeOk returns a tuple with the ActivatedStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedStake

`func (o *ValidatorValidator) SetActivatedStake(v int32)`

SetActivatedStake sets ActivatedStake field to given value.

### HasActivatedStake

`func (o *ValidatorValidator) HasActivatedStake() bool`

HasActivatedStake returns a boolean if a field has been set.

### GetStakePercentage

`func (o *ValidatorValidator) GetStakePercentage() int32`

GetStakePercentage returns the StakePercentage field if non-nil, zero value otherwise.

### GetStakePercentageOk

`func (o *ValidatorValidator) GetStakePercentageOk() (*int32, bool)`

GetStakePercentageOk returns a tuple with the StakePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakePercentage

`func (o *ValidatorValidator) SetStakePercentage(v int32)`

SetStakePercentage sets StakePercentage field to given value.

### HasStakePercentage

`func (o *ValidatorValidator) HasStakePercentage() bool`

HasStakePercentage returns a boolean if a field has been set.

### GetCommission

`func (o *ValidatorValidator) GetCommission() int32`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *ValidatorValidator) GetCommissionOk() (*int32, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *ValidatorValidator) SetCommission(v int32)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *ValidatorValidator) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetEpochCredits

`func (o *ValidatorValidator) GetEpochCredits() []int32`

GetEpochCredits returns the EpochCredits field if non-nil, zero value otherwise.

### GetEpochCreditsOk

`func (o *ValidatorValidator) GetEpochCreditsOk() (*[]int32, bool)`

GetEpochCreditsOk returns a tuple with the EpochCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochCredits

`func (o *ValidatorValidator) SetEpochCredits(v []int32)`

SetEpochCredits sets EpochCredits field to given value.

### HasEpochCredits

`func (o *ValidatorValidator) HasEpochCredits() bool`

HasEpochCredits returns a boolean if a field has been set.

### GetEpochVoteAccount

`func (o *ValidatorValidator) GetEpochVoteAccount() bool`

GetEpochVoteAccount returns the EpochVoteAccount field if non-nil, zero value otherwise.

### GetEpochVoteAccountOk

`func (o *ValidatorValidator) GetEpochVoteAccountOk() (*bool, bool)`

GetEpochVoteAccountOk returns a tuple with the EpochVoteAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochVoteAccount

`func (o *ValidatorValidator) SetEpochVoteAccount(v bool)`

SetEpochVoteAccount sets EpochVoteAccount field to given value.

### HasEpochVoteAccount

`func (o *ValidatorValidator) HasEpochVoteAccount() bool`

HasEpochVoteAccount returns a boolean if a field has been set.

### GetLastVote

`func (o *ValidatorValidator) GetLastVote() int32`

GetLastVote returns the LastVote field if non-nil, zero value otherwise.

### GetLastVoteOk

`func (o *ValidatorValidator) GetLastVoteOk() (*int32, bool)`

GetLastVoteOk returns a tuple with the LastVote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVote

`func (o *ValidatorValidator) SetLastVote(v int32)`

SetLastVote sets LastVote field to given value.

### HasLastVote

`func (o *ValidatorValidator) HasLastVote() bool`

HasLastVote returns a boolean if a field has been set.

### GetNodePubkey

`func (o *ValidatorValidator) GetNodePubkey() string`

GetNodePubkey returns the NodePubkey field if non-nil, zero value otherwise.

### GetNodePubkeyOk

`func (o *ValidatorValidator) GetNodePubkeyOk() (*string, bool)`

GetNodePubkeyOk returns a tuple with the NodePubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePubkey

`func (o *ValidatorValidator) SetNodePubkey(v string)`

SetNodePubkey sets NodePubkey field to given value.

### HasNodePubkey

`func (o *ValidatorValidator) HasNodePubkey() bool`

HasNodePubkey returns a boolean if a field has been set.

### GetRootSlot

`func (o *ValidatorValidator) GetRootSlot() int32`

GetRootSlot returns the RootSlot field if non-nil, zero value otherwise.

### GetRootSlotOk

`func (o *ValidatorValidator) GetRootSlotOk() (*int32, bool)`

GetRootSlotOk returns a tuple with the RootSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSlot

`func (o *ValidatorValidator) SetRootSlot(v int32)`

SetRootSlot sets RootSlot field to given value.

### HasRootSlot

`func (o *ValidatorValidator) HasRootSlot() bool`

HasRootSlot returns a boolean if a field has been set.

### GetVotePubkey

`func (o *ValidatorValidator) GetVotePubkey() string`

GetVotePubkey returns the VotePubkey field if non-nil, zero value otherwise.

### GetVotePubkeyOk

`func (o *ValidatorValidator) GetVotePubkeyOk() (*string, bool)`

GetVotePubkeyOk returns a tuple with the VotePubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotePubkey

`func (o *ValidatorValidator) SetVotePubkey(v string)`

SetVotePubkey sets VotePubkey field to given value.

### HasVotePubkey

`func (o *ValidatorValidator) HasVotePubkey() bool`

HasVotePubkey returns a boolean if a field has been set.

### GetBlockProduction

`func (o *ValidatorValidator) GetBlockProduction() ValidatorValidatorBlockProduction`

GetBlockProduction returns the BlockProduction field if non-nil, zero value otherwise.

### GetBlockProductionOk

`func (o *ValidatorValidator) GetBlockProductionOk() (*ValidatorValidatorBlockProduction, bool)`

GetBlockProductionOk returns a tuple with the BlockProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockProduction

`func (o *ValidatorValidator) SetBlockProduction(v ValidatorValidatorBlockProduction)`

SetBlockProduction sets BlockProduction field to given value.

### HasBlockProduction

`func (o *ValidatorValidator) HasBlockProduction() bool`

HasBlockProduction returns a boolean if a field has been set.

### GetDelegatingStakeAccounts

`func (o *ValidatorValidator) GetDelegatingStakeAccounts() []StakeAccount`

GetDelegatingStakeAccounts returns the DelegatingStakeAccounts field if non-nil, zero value otherwise.

### GetDelegatingStakeAccountsOk

`func (o *ValidatorValidator) GetDelegatingStakeAccountsOk() (*[]StakeAccount, bool)`

GetDelegatingStakeAccountsOk returns a tuple with the DelegatingStakeAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatingStakeAccounts

`func (o *ValidatorValidator) SetDelegatingStakeAccounts(v []StakeAccount)`

SetDelegatingStakeAccounts sets DelegatingStakeAccounts field to given value.

### HasDelegatingStakeAccounts

`func (o *ValidatorValidator) HasDelegatingStakeAccounts() bool`

HasDelegatingStakeAccounts returns a boolean if a field has been set.

### GetDelegatorCount

`func (o *ValidatorValidator) GetDelegatorCount() int32`

GetDelegatorCount returns the DelegatorCount field if non-nil, zero value otherwise.

### GetDelegatorCountOk

`func (o *ValidatorValidator) GetDelegatorCountOk() (*int32, bool)`

GetDelegatorCountOk returns a tuple with the DelegatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatorCount

`func (o *ValidatorValidator) SetDelegatorCount(v int32)`

SetDelegatorCount sets DelegatorCount field to given value.

### HasDelegatorCount

`func (o *ValidatorValidator) HasDelegatorCount() bool`

HasDelegatorCount returns a boolean if a field has been set.

### GetLocation

`func (o *ValidatorValidator) GetLocation() NonvalidatorsLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ValidatorValidator) GetLocationOk() (*NonvalidatorsLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ValidatorValidator) SetLocation(v NonvalidatorsLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ValidatorValidator) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMoniker

`func (o *ValidatorValidator) GetMoniker() string`

GetMoniker returns the Moniker field if non-nil, zero value otherwise.

### GetMonikerOk

`func (o *ValidatorValidator) GetMonikerOk() (*string, bool)`

GetMonikerOk returns a tuple with the Moniker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoniker

`func (o *ValidatorValidator) SetMoniker(v string)`

SetMoniker sets Moniker field to given value.

### HasMoniker

`func (o *ValidatorValidator) HasMoniker() bool`

HasMoniker returns a boolean if a field has been set.

### GetWebsite

`func (o *ValidatorValidator) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ValidatorValidator) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ValidatorValidator) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ValidatorValidator) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetPictureURL

`func (o *ValidatorValidator) GetPictureURL() string`

GetPictureURL returns the PictureURL field if non-nil, zero value otherwise.

### GetPictureURLOk

`func (o *ValidatorValidator) GetPictureURLOk() (*string, bool)`

GetPictureURLOk returns a tuple with the PictureURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureURL

`func (o *ValidatorValidator) SetPictureURL(v string)`

SetPictureURL sets PictureURL field to given value.

### HasPictureURL

`func (o *ValidatorValidator) HasPictureURL() bool`

HasPictureURL returns a boolean if a field has been set.

### GetVersion

`func (o *ValidatorValidator) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ValidatorValidator) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ValidatorValidator) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ValidatorValidator) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDetails

`func (o *ValidatorValidator) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ValidatorValidator) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ValidatorValidator) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ValidatorValidator) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetAsn

`func (o *ValidatorValidator) GetAsn() ValidatorValidatorAsn`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ValidatorValidator) GetAsnOk() (*ValidatorValidatorAsn, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ValidatorValidator) SetAsn(v ValidatorValidatorAsn)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *ValidatorValidator) HasAsn() bool`

HasAsn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


